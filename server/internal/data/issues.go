package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/lucabrx/wuhu/internal/validator"
)

var (
	BacklogStatus   = "Backlog"
	TodoStatus      = "Todo"
	InProgress      = "In Progress"
	BlockedStatus   = "Blocked"
	DoneStatus      = "Done"
	CancelledStatus = "Cancelled"
	CheckStatus     = "Check"
)

type IssueModal struct {
	DB *sql.DB
}

type Issue struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Status      string    `json:"status"`
	Priority    string    `json:"priority,omitempty"`
	DueDate     time.Time `json:"dueDate,omitempty"`
	Version     int64     `json:"version"`
}

type UserWorkspaceIssue struct {
	UserId      int64 `json:"userId"`
	WorkspaceId int64 `json:"workspaceId"`
	IssueId     int64 `json:"issueId"`
	CreatorId   int64 `json:"creatorId"`
}

func ValidateIssue(v *validator.Validator, issue *Issue) {
	v.Check(issue.Title != "", "title", "title must be provided.")
	v.Check(issue.Status != "", "status", "status must be provided.")
	v.Check(validator.PermittedValue(issue.Status, BacklogStatus, TodoStatus, InProgress, BlockedStatus, DoneStatus, CancelledStatus, CheckStatus), "status", "status is not valid.")
}

func (m *IssueModal) Insert(issue *Issue) error {
	query := `
        INSERT INTO issues (title, description, status, priority, due_date)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, created_at, version
    `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	args := []interface{}{
		issue.Title,
		newNullString(issue.Description),
		issue.Status,
		newNullString(issue.Priority),
		newNullDate(issue.DueDate),
	}

	if err := m.DB.QueryRowContext(ctx, query, args...).Scan(&issue.ID, &issue.CreatedAt, &issue.Version); err != nil {
		return err
	}

	return nil
}

func (m *IssueModal) Get(id int64) (*Issue, error) {
	query := `SELECT id, created_at, title, COALESCE(description, ''), status, COALESCE(priority, ''), COALESCE(due_date, '0001-01-01'), version FROM issues WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var issue Issue
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&issue.ID,
		&issue.CreatedAt,
		&issue.Title,
		&issue.Description,
		&issue.Status,
		&issue.Priority,
		&issue.DueDate,
		&issue.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	if issue.DueDate == time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC) {
		issue.DueDate = time.Time{}
	}

	return &issue, nil
}

func (m *IssueModal) Update(issue *Issue) error {
	query := `
        UPDATE issues
        SET title = $1, description = $2, status = $3, priority = $4, due_date = $5, version = version + 1
        WHERE id = $6 AND version = $7
        RETURNING version
    `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{
		issue.Title,
		newNullString(issue.Description),
		issue.Status,
		newNullString(issue.Priority),
		newNullDate(issue.DueDate),
		issue.ID,
		issue.Version,
	}

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&issue.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil
}

func (m *IssueModal) Delete(id int64) error {
	query := `DELETE FROM issues WHERE id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *IssueModal) InsertUserWorkspaceIssue(userWorkspaceIssue *UserWorkspaceIssue) error {
	fmt.Println("UserWorkspaceIssue: ", userWorkspaceIssue.CreatorId)
	query := `
        INSERT INTO user_workspace_issues (issue_id, creator_id, user_id, workspace_id)
        VALUES ($1, $2, $3, $4)
    `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := make([]interface{}, 4)
	args[0] = userWorkspaceIssue.IssueId
	args[1] = userWorkspaceIssue.CreatorId

	if userWorkspaceIssue.UserId == 0 {
		args[2] = nil
	} else {
		args[2] = userWorkspaceIssue.UserId
	}

	if userWorkspaceIssue.WorkspaceId == 0 {
		args[3] = nil
	} else {
		args[3] = userWorkspaceIssue.WorkspaceId
	}

	fmt.Println("Executing SQL:", query)
	fmt.Println("With Parameters:", args)

	_, err := m.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
