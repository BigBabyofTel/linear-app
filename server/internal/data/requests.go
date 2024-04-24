package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type RequestModal struct {
	DB *sql.DB
}

type Request struct {
	UserId      int64  `json:"userId"`
	WorkspaceId int64  `json:"workspaceId"`
	Name        string `json:"name"`
	ImageUrl    string `json:"imageUrl"`
}

func (m *RequestModal) Insert(request *Request) error {
	query := `INSERT INTO requests (user_id, workspace_id)
			  VALUES ($1, $2)
			`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, request.UserId, request.WorkspaceId)
	return err
}

func (m *RequestModal) Delete(userId int64, workspaceId int64) error {
	query := `
			DELETE FROM requests
			WHERE user_id = $1 AND workspace_id = $2 
			`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, userId, workspaceId)
	if err != nil {
		return err
	}
	rowsEffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsEffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (m *RequestModal) Get(userId int64, workspaceId int64) (*Request, error) {
	query := `
            SELECT user_id, workspace_id
            FROM requests
            WHERE user_id = $1 AND workspace_id = $2
            `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var request Request
	err := m.DB.QueryRowContext(ctx, query, userId, workspaceId).Scan(
		&request.UserId,
		&request.WorkspaceId,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &request, nil
}

func (m *RequestModal) GetAll(userId int64) ([]*Request, error) {
	query := `
			SELECT r.user_id, r.workspace_id, COALESCE(w.image, ''), w.name
			FROM requests r
			INNER JOIN workspaces w ON w.id = r.workspace_id
			WHERE user_id = $1
			`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []*Request

	for rows.Next() {
		var request Request
		err := rows.Scan(&request.UserId, &request.WorkspaceId, &request.ImageUrl, &request.Name)
		if err != nil {
			return nil, err
		}
		requests = append(requests, &request)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return requests, nil
}
