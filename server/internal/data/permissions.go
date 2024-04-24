package data

import (
	"context"
	"database/sql"
	"time"
)

type PermissionModal struct {
	DB *sql.DB
}

type Permissions string

var (
	AdminPermission     = "workspace:admin"
	ModeratorPermission = "workspace:moderator"
	MemberPermission    = "workspace:member"
)

// if code == "workspace:member" --> allow member,moderator,admin
// if code == "workspace:moderator" --> allow moderator,admin
// if code == "workspace:admin" --> allow admin

func (p Permissions) Contains(code string) bool {
	if code == "workspace:member" {
		if p == "workspace:member" || p == "workspace:moderator" || p == "workspace:admin" {
			return true
		}
	}

	if code == "workspace:moderator" {
		if p == "workspace:moderator" || p == "workspace:admin" {
			return true
		}
	}

	if code == "workspace:admin" {
		if p == "workspace:admin" {
			return true
		}
	}

	return false
}

func (p *PermissionModal) GetForUser(userId int64, workspaceId int64) (Permissions, error) {
	query := `SELECT p.code FROM permissions p
			  JOIN user_workspace_permissions uwp ON uwp.permission_id = p.id
			  WHERE uwp.user_id = $1 AND uwp.workspace_id = $2
			`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var permissions Permissions
	err := p.DB.QueryRowContext(ctx, query, userId, workspaceId).Scan(&permissions)
	if err != nil {
		return "", err
	}

	return permissions, nil
}

func (m *PermissionModal) AddForUser(userId int64, workspaceId int64, code string) error {
	query := `INSERT INTO user_workspace_permissions (user_id, workspace_id, permission_id)
			  VALUES ($1, $2, (SELECT id FROM permissions WHERE code = $3))
			`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, userId, workspaceId, code)
	return err
}

func (m *PermissionModal) RemoveForUser(userId int64, workspaceId int64) error {
	query := ` DELETE FROM user_workspace_permissions 
			   WHERE user_id = $1 AND workspace_id = $2
			 `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.ExecContext(ctx, query, userId, workspaceId)
	if err != nil {
		return err
	}
	if rows, _ := rows.RowsAffected(); rows == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (m *PermissionModal) UpdateForUser(userId int64, workspaceId int64, permissions Permissions) error {
	query := `UPDATE user_workspace_permissions
			  SET permission_id = (SELECT id FROM permissions WHERE code = $3)
			  WHERE user_id = $1 AND workspace_id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, userId, workspaceId, permissions)

	return err
}
