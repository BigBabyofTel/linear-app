package api

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lucabrx/wuhu/internal/data"
)

func (a *app) inviteUserToWorkspaceHandler(w http.ResponseWriter, r *http.Request) {
	session := a.contextGetUser(r)
	slug := chi.URLParam(r, "slug")
	workspace, err := a.DB.Workspace.Get(0, slug, session.ID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundResponse(w, r)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return

	}
	var input struct {
		Email string `json:"email"`
	}

	if err = a.readJSON(w, r, &input); err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	user, err := a.DB.User.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundEmailResponse(w, r)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	request := &data.Request{
		WorkspaceId: workspace.ID,
		UserId:      user.ID,
	}
	if err = a.DB.Request.Insert(request); err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
	if err = a.writeJSON(w, http.StatusCreated, envelope{"message": "success"}, nil); err != nil {
		a.serverErrorResponse(w, r, err)
	}

}

func (a *app) answerRequestHandler(w http.ResponseWriter, r *http.Request) {
	session := a.contextGetUser(r)

	var input struct {
		WorkspaceId int64 `json:"WorkspaceId"`
		Answer      bool  `json:"answer"`
	}
	if err := a.readJSON(w, r, &input); err != nil {
		a.badRequestResponse(w, r, err)
		return
	}
	request, err := a.DB.Request.Get(session.ID, input.WorkspaceId)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundResponse(w, r)
		default:
			a.serverErrorResponse(w, r, err)
		}
	}

	var message string
	var value bool
	if input.Answer {
		err = a.DB.Permission.AddForUser(session.ID, input.WorkspaceId, data.MemberPermission)
		if err != nil {
			a.serverErrorResponse(w, r, err)
			return
		}
		message = "Welcome to your new Workspace"
		value = true
	} else {
		message = "Well, maybe next time 🥲"
		value = false
	}

	err = a.DB.Request.Delete(request.UserId, request.WorkspaceId)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	if err = a.writeJSON(w, http.StatusOK, envelope{"message": message, "value": value}, nil); err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *app) getAllRequestsHandler(w http.ResponseWriter, r *http.Request) {
	session := a.contextGetUser(r)

	requests, err := a.DB.Request.GetAll(session.ID)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	if err = a.writeJSON(w, http.StatusOK, envelope{"requests": requests}, nil); err != nil {
		a.serverErrorResponse(w, r, err)
	}
}
