package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/lucabrx/wuhu/internal/data"
	"github.com/lucabrx/wuhu/internal/validator"
)

func (a *app) createIssueHandler(w http.ResponseWriter, r *http.Request) {
	session := a.contextGetUser(r)
	fmt.Println("Session: ", session.ID)
	var input struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Status      string    `json:"status"`
		Priority    string    `json:"priority"`
		DueDate     time.Time `json:"dueDate"`
		UserId      int64     `json:"userId"`
		WorkspaceId int64     `json:"workspaceId"`
	}

	if err := a.readJSON(w, r, &input); err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	issue := &data.Issue{
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		Priority:    input.Priority,
		DueDate:     input.DueDate,
	}

	v := validator.New()
	if data.ValidateIssue(v, issue); !v.Valid() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	if err := a.DB.Issue.Insert(issue); err != nil {
		fmt.Println("Error inserting issue: ", err)
		a.serverErrorResponse(w, r, err)
		return
	}

	userWorkspaceIssue := &data.UserWorkspaceIssue{
		UserId:      input.UserId,
		WorkspaceId: input.WorkspaceId,
		IssueId:     issue.ID,
		CreatorId:   session.ID,
	}

	if err := a.DB.Issue.InsertUserWorkspaceIssue(userWorkspaceIssue); err != nil {
		fmt.Println("Error inserting user workspace issue: ", err)
		a.serverErrorResponse(w, r, err)
		return
	}

	if err := a.writeJSON(w, http.StatusCreated, envelope{"issue": issue}, nil); err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *app) getAllWorkspaceIssuesHandler(w http.ResponseWriter, r *http.Request) {
	params := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(params, 10, 64)

	issues, err := a.DB.Issue.GetIssuesByWorkspace(id)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	if err := a.writeJSON(w, http.StatusOK, envelope{"issues": issues}, nil); err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *app) getAllUserIssuesHandler(w http.ResponseWriter, r *http.Request) {
	session := a.contextGetUser(r)

	var meta struct {
		Search string
		data.Filters
	}

	qs := r.URL.Query()
	v := validator.New()

	meta.Filters.Sort = a.readString(qs, "sort", "-i.id")
	meta.Filters.SortSafelist = []string{"i.id", "i.title", "i.status", "i.priority", "i.due_date", "i.created_at", "-i.id", "-i.title", "-i.status", "-i.priority", "-i.due_date", "-i.created_at"}
	meta.Filters.Page = a.readInt(qs, "page", 1, v)
	meta.Filters.PageSize = a.readInt(qs, "pageSize", 20, v)
	meta.Search = a.readString(qs, "search", "")

	if data.ValidateFilters(v, meta.Filters); !v.Valid() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	issues, metadata, err := a.DB.Issue.GetIssuesByUser(session.ID, meta.Search, meta.Filters)
	fmt.Println("Error: ", err)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundResponse(w, r)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	if err := a.writeJSON(w, http.StatusOK, envelope{"metadata": metadata, "issues": issues}, nil); err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *app) getIssueHandler(w http.ResponseWriter, r *http.Request) {
	session := a.contextGetUser(r)
	params := chi.URLParam(r, "issueId")
	id, _ := strconv.ParseInt(params, 10, 64)

	issue, err := a.DB.Issue.GetIssueByUser(session.ID, id)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	if err := a.writeJSON(w, http.StatusOK, envelope{"issue": issue}, nil); err != nil {
		a.serverErrorResponse(w, r, err)
	}
}
