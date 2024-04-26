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

func (m *IssueModal) GetIssuesByWorkspace(workspaceId int64) ([]*Issue, error) {
	query := `SELECT i.id, i.created_at, i.title, COALESCE(i.description, ''), i.status, COALESCE(i.priority, ''), COALESCE(i.due_date, '0001-01-01'), i.version FROM issues i JOIN user_workspace_issues uwi ON i.id = uwi.issue_id WHERE uwi.workspace_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, workspaceId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var issues []*Issue
	for rows.Next() {
		var issue Issue
		err := rows.Scan(
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
			return nil, err
		}

		if issue.DueDate == time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC) {
			issue.DueDate = time.Time{}
		}

		issues = append(issues, &issue)
	}

	return issues, nil
}

func (m *IssueModal) GetIssuesByUser(userId int64, search string, filters Filters) ([]*Issue, Metadata, error) {
	baseQuery := `SELECT count(*) OVER(), i.id, i.created_at, i.title, i.status, COALESCE(i.priority, ''), COALESCE(i.due_date, '0001-01-01'),
				  i.version 
				  FROM issues i 
				  JOIN user_workspace_issues uwi 
				  ON i.id = uwi.issue_id 
				  WHERE uwi.creator_id = $1 
				  `

	if search != "" {
		// only title search
		baseQuery += `AND (LOWER(i.title) LIKE '%' || LOWER($4) || '%')`
	}

	advQuery := fmt.Sprintf(` ORDER BY %s %s, i.id ASC LIMIT $2 OFFSET $3`,
		filters.sortColumn(), filters.sortDirection(),
	)

	args := []interface{}{userId, filters.limit(), filters.offset()}
	if search != "" {
		args = append(args, search)
	}
	query := baseQuery + advQuery
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, args...)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, Metadata{}, ErrRecordNotFound
		default:
			return nil, Metadata{}, err
		}
	}
	defer rows.Close()

	totalRecords := 0
	var issues []*Issue
	for rows.Next() {
		var issue Issue
		err := rows.Scan(
			&totalRecords,
			&issue.ID,
			&issue.CreatedAt,
			&issue.Title,
			&issue.Status,
			&issue.Priority,
			&issue.DueDate,
			&issue.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		if issue.DueDate == time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC) {
			issue.DueDate = time.Time{}
		}

		issues = append(issues, &issue)
	}

	if totalRecords == 0 {
		return nil, Metadata{}, ErrRecordNotFound
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return issues, metadata, nil
}

func (m *IssueModal) GetIssueByUser(userId, issueId int64) (*Issue, error) {
	query := `SELECT i.id, i.created_at, i.title, COALESCE(i.description, ''), i.status, COALESCE(i.priority, ''), COALESCE(i.due_date, '0001-01-01'), i.version FROM issues i JOIN user_workspace_issues uwi ON i.id = uwi.issue_id WHERE uwi.user_id = $1 AND i.id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var issue Issue
	err := m.DB.QueryRowContext(ctx, query, userId, issueId).Scan(
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
		return nil, err
	}

	if issue.DueDate == time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC) {
		issue.DueDate = time.Time{}
	}

	return &issue, nil
}

func (m *IssueModal) GetIssueByWorkspace(workspaceId, issueId int64) (*Issue, error) {
	query := `SELECT i.id, i.created_at, i.title, COALESCE(i.description, ''), i.status, COALESCE(i.priority, ''), COALESCE(i.due_date, '0001-01-01'), i.version FROM issues i JOIN user_workspace_issues uwi ON i.id = uwi.issue_id WHERE uwi.workspace_id = $1 AND i.id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var issue Issue
	err := m.DB.QueryRowContext(ctx, query, workspaceId, issueId).Scan(
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
		return nil, err
	}

	if issue.DueDate == time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC) {
		issue.DueDate = time.Time{}
	}

	return &issue, nil
}
