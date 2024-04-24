package data

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base32"
	"github.com/lucabrx/wuhu/internal/random"
	"math/rand"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/require"
)

type TestDB struct {
	User  UserModel
	Token TokenModel
}

var testDB *TestDB

func TestMain(m *testing.M) {
	db, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5431/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	testDB = &TestDB{
		User:  UserModel{DB: db},
		Token: TokenModel{DB: db},
	}

	m.Run()
}

func createRandomUser(t *testing.T) User {
	user := &User{
		Name:  random.RandString(6),
		Email: random.RandString(6) + "@gmail.com",
	}

	err := user.Password.Set(random.RandString(10))
	user.Password.plaintext = nil
	require.NoError(t, err)
	err = testDB.User.Insert(user)
	require.NoError(t, err)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return *user
}

func createRandomUserToken(t *testing.T, user User) Token {
	token := &Token{
		UserID: user.ID,
		Expiry: time.Now().Add(24 * time.Hour),
		Scope:  ScopeAuthentication,
	}
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	require.NoError(t, err)
	token.Plaintext = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)
	hash := sha256.Sum256([]byte(token.Plaintext))
	token.Hash = hash[:]

	err = testDB.Token.Insert(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	return *token
}
