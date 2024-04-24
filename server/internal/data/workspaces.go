package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lucabrx/wuhu/internal/validator"
)

var ErrDuplicateSlug = errors.New("duplicate slug")

type WorkspaceModal struct {
	DB *sql.DB
}

type Workspace struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	Image     string    `json:"image,omitempty"`
	Slug      string    `json:"slug"`
	Version   int64     `json:"version"`
	RoleId    int64     `json:"roleId"`
	Role      string    `json:"role"`
}

func ValidateWorkspace(v *validator.Validator, name string) {
	v.Check(name != "", "name", "Workspace name must be provided")
	v.Check(len(name) >= 3, "name", "Workspace name must be longer then 3 characters.")
}

func (m *WorkspaceModal) Insert(workspace *Workspace) error {
	query := `INSERT INTO workspaces (name, image, slug) VALUES ($1, $2, $3) RETURNING id, created_at, version`

	args := []interface{}{workspace.Name, workspace.Image, workspace.Slug}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&workspace.ID, &workspace.CreatedAt, &workspace.Version)
	if err != nil {
		switch {
		case err.Error() == `ERROR: duplicate key value violates unique constraint "workspaces_slug_key" (SQLSTATE 23505)`:
			return ErrDuplicateSlug
		default:
			return err
		}

	}
	return nil
}

func (w WorkspaceModal) Get(id int64, slug string, userId int64) (*Workspace, error) {
	query := `SELECT w.id, w.created_at, w.name, w.slug, COALESCE(w.image, ''), w.version, uwp.permission_id
			  FROM workspaces w
			  INNER JOIN user_workspace_permissions uwp ON uwp.workspace_id = w.id
			  INNER JOIN users u ON uwp.user_id = $3
			  WHERE w.id = $1 OR w.slug = $2
              `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var workspace Workspace
	err := w.DB.QueryRowContext(ctx, query, id, slug, userId).Scan(
		&workspace.ID,
		&workspace.CreatedAt,
		&workspace.Name,
		&workspace.Slug,
		&workspace.Image,
		&workspace.Version,
		&workspace.RoleId,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &workspace, nil
}

func (m *WorkspaceModal) Update(workspace *Workspace) error {
	query := `UPDATE workspaces SET name = $1, image = $2, slug = $3, version = version + 1 WHERE id = $4 RETURNING version`
	args := []interface{}{workspace.Name, workspace.Image, workspace.Slug, workspace.ID}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&workspace.Version)
	if err != nil {
		switch {
		case err.Error() == `ERROR: duplicate key value violates unique constraint "workspaces_slug_key" (SQLSTATE 23505)`:
			return ErrDuplicateSlug
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (m *WorkspaceModal) Delete(slug string) error {
	query := `DELETE FROM workspaces WHERE slug = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, slug)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (m *WorkspaceModal) GetAllWorkspaceUsers(workspaceId int64) ([]*User, error) {
	query := `SELECT count(*) OVER(), u.id, COALESCE(u.name), COALESCE(u.email), u.created_at, COALESCE(u.image, ''), uwp.permission_id
 				  FROM user_workspace_permissions uwp
 				  INNER JOIN users u ON uwp.user_id = u.id
                  WHERE uwp.workspace_id = $1 
    `

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, workspaceId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	totalRecords := 0
	var users []*User
	for rows.Next() {
		user := &User{}
		err := rows.Scan(
			&totalRecords,
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
			&user.Image,
			&user.RoleId,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (m *WorkspaceModal) GetUserWorkspaces(userId int64) ([]*Workspace, error) {
	query := `SELECT DISTINCT w.id, w.created_at, w.name, w.slug, COALESCE(w.IMAGE, ''), w.version, uwp.permission_id
			  FROM workspaces w
			  INNER JOIN user_workspace_permissions uwp ON uwp.workspace_id = w.id
			  INNER JOIN users u ON uwp.user_id = $1
			  `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workspaces []*Workspace
	for rows.Next() {
		workspace := &Workspace{}
		err := rows.Scan(
			&workspace.ID,
			&workspace.CreatedAt,
			&workspace.Name,
			&workspace.Slug,

			&workspace.Image,
			&workspace.Version,
			&workspace.RoleId,
		)
		if err != nil {
			return nil, err
		}
		workspaces = append(workspaces, workspace)
	}

	return workspaces, nil
}

func (m *WorkspaceModal) RemoveWorkspaceImage(slug string) error {
	query := `UPDATE workspaces SET image = NULL WHERE slug = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.ExecContext(ctx, query, slug)
	if err != nil {
		return err
	}

	if rows, _ := rows.RowsAffected(); rows == 0 {
		return ErrRecordNotFound
	}

	return nil
}
