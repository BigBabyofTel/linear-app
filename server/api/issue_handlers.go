package api

import (
	"fmt"
	"net/http"
	"time"

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
