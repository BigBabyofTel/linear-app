package data

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"time"

	"github.com/lucabrx/wuhu/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

var ErrDuplicateEmail = errors.New("duplicate email")
var AnonymousUser = &User{}

type UserModel struct {
	DB *sql.DB
}
type User struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	GithubId  int64     `json:"githubId,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Image     string    `json:"image,omitempty"`
	Password  password  `json:"-"`
	Verified  bool      `json:"verified"`
	RoleId    int64     `json:"roleId,omitempty"`
	Role      string    `json:"role,omitempty"`
}

type password struct {
	plaintext *string
	hash      []byte
}

func (u *User) IsAnonymous() bool {
	return u == AnonymousUser
}

func (p *password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}

	p.plaintext = &plaintextPassword
	p.hash = hash

	return nil
}

func (p *password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}

func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(password) <= 72, "password", "must not be more than 72 bytes long")
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Name != "", "name", "must be provided")
	v.Check(len(user.Name) <= 500, "name", "must not be more than 500 bytes long")

	ValidateEmail(v, user.Email)

	if user.Password.plaintext != nil {
		ValidatePasswordPlaintext(v, *user.Password.plaintext)
	}

	if user.Password.hash == nil {
		panic("missing password hash for user")
	}
}

func (m *UserModel) Insert(user *User) error {
	query := `INSERT INTO users (github_id, name, email, image, password, verified) 
              VALUES ($1, $2, $3, $4, $5, $6)
              RETURNING id, created_at`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{
		newNullInt64(user.GithubId),
		newNullString(user.Name),
		newNullString(user.Email),
		newNullString(user.Image),
		newNullPassword(user.Password),
		newNullBool(user.Verified),
	}

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		switch {
		case err.Error() == `ERROR: duplicate key value violates unique constraint "users_email_key" (SQLSTATE 23505)`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}

	return nil
}

func (m *UserModel) GetByID(id int64) (*User, error) {
	query := "SELECT id, created_at, COALESCE(name, ''), COALESCE(email, ''), COALESCE(image, ''), COALESCE(password, ''), COALESCE(github_id, 0), verified FROM users WHERE id = $1"
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user User
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.Image,
		&user.Password.hash,
		&user.GithubId,
		&user.Verified,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (m *UserModel) GetByEmail(email string) (*User, error) {
	query := "SELECT id, created_at, COALESCE(name, ''), COALESCE(email, ''), COALESCE(image, ''), COALESCE(password, ''), COALESCE(github_id, 0), verified FROM users WHERE email = $1"
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user User
	err := m.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.Image,
		&user.Password.hash,
		&user.GithubId,
		&user.Verified,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (m *UserModel) GetForToken(tokenScope, tokenPlaintext string) (*User, error) {
	tokenHash := sha256.Sum256([]byte(tokenPlaintext))
	query := `SELECT u.id, u.created_at, COALESCE(u.name, ''), COALESCE(u.email, '') ,COALESCE(u.image,''), COALESCE(u.github_id, 0),
			  COALESCE(u.password, ''), verified
        	  FROM users u
        	  INNER JOIN tokens t
        	  ON u.id = t.user_id
        	  WHERE t.hash = $1
        	  AND t.scope = $2
        	  AND t.expiry > $3`

	args := []any{tokenHash[:], tokenScope, time.Now()}

	var user User
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.Image,
		&user.GithubId,
		&user.Password.hash,
		&user.Verified,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (m *UserModel) Delete(id int64) error {
	query := "DELETE FROM users WHERE id = $1"
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *UserModel) Update(user *User) error {
	query := `UPDATE users SET name = $2, email = $3, image = $4, password = $5, verified = $6 WHERE id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{
		user.ID,
		newNullString(user.Name),
		newNullString(user.Email),
		newNullString(user.Image),
		newNullPassword(user.Password),
		newNullBool(user.Verified),
	}

	_, err := m.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
