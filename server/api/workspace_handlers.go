package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/lucabrx/wuhu/internal/data"
	"github.com/lucabrx/wuhu/internal/random"
	"github.com/lucabrx/wuhu/internal/validator"
)

func (a *app) createWorkspaceHandler(w http.ResponseWriter, r *http.Request) {
	session := a.contextGetUser(r)

	var input struct {
		Name  string `json:"name"`
		Slug  string `json:"slug"`
		Image string `json:"image"`
	}

	if err := a.readJSON(w, r, &input); err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	if input.Slug == "" {
		randNum := random.RandInt(1, 9999)
		input.Slug = input.Name + "-" + fmt.Sprint(randNum)
	}

	v := validator.New()
	if data.ValidateWorkspace(v, input.Name); !v.Valid() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	workspace := &data.Workspace{
		Name:  input.Name,
		Slug:  input.Slug,
		Image: input.Image,
	}

	if err := a.DB.Workspace.Insert(workspace); err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateSlug):
			v.AddError("slug", "a workspace with this slug already exists")
			a.failedValidationResponse(w, r, v.Errors)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	if err := a.DB.Permission.AddForUser(session.ID, workspace.ID, data.AdminPermission); err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	workspace.Role = data.AdminPermission
	workspace.RoleId = 1

	if err := a.writeJSON(w, http.StatusCreated, envelope{"workspace": workspace}, nil); err != nil {
		a.serverErrorResponse(w, r, err)
	}

}

func (a *app) getWorkspaceHandler(w http.ResponseWriter, r *http.Request) {
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
	switch workspace.RoleId {
	case 1:
		workspace.Role = "admin"
	case 2:
		workspace.Role = "moderator"
	case 3:
		workspace.Role = "member"
	}

	if err := a.writeJSON(w, http.StatusOK, envelope{"workspace": workspace}, nil); err != nil {
		a.serverErrorResponse(w, r, err)
	}

}

func (a *app) deleteWorkspaceHandler(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	if err := a.DB.Workspace.Delete(slug); err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundResponse(w, r)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	if err := a.writeJSON(w, http.StatusOK, envelope{"message": "workspace deleted successfully"}, nil); err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *app) updateWorkspaceHandler(w http.ResponseWriter, r *http.Request) {
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
		Name string `json:"name"`
		Slug string `json:"slug"`
	}
	if err = a.readJSON(w, r, &input); err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
	workspace.Name = input.Name
	workspace.Slug = input.Slug
	v := validator.New()
	data.ValidateWorkspace(v, workspace.Name)
	if !v.Valid() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}
	if err = a.DB.Workspace.Update(workspace); err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateSlug):
			v.AddError("slug", "a workspace with this slug already exists")
			a.failedValidationResponse(w, r, v.Errors)
		case errors.Is(err, data.ErrEditConflict):
			a.editConflictResponse(w, r)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}
	updateWorkspace, err := a.DB.Workspace.Get(0, workspace.Slug, session.ID)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	if err = a.writeJSON(w, http.StatusOK, envelope{"workspace": updateWorkspace}, nil); err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *app) removeWorkspaceImageHandler(w http.ResponseWriter, r *http.Request) {
	session := a.contextGetUser(r)
	slug := chi.URLParam(r, "slug")

	if err := a.DB.Workspace.RemoveWorkspaceImage(slug); err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
	workspace, err := a.DB.Workspace.Get(0, slug, session.ID)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	if err = a.writeJSON(w, http.StatusOK, envelope{"workspace": workspace}, nil); err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *app) getAllUserWorkspacesHandler(w http.ResponseWriter, r *http.Request) {
	session := a.contextGetUser(r)

	workspaces, err := a.DB.Workspace.GetUserWorkspaces(session.ID)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	if err = a.writeJSON(w, http.StatusOK, envelope{"workspaces": workspaces}, nil); err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *app) changeUserRoleHandler(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	fmt.Println(slug)
	userIdParams := chi.URLParam(r, "userId")
	workspacesIdParams := chi.URLParam(r, "id")

	userId, err := strconv.ParseInt(userIdParams, 10, 64)
	if userId <= 0 {
		a.notFoundResponse(w, r)
		return
	}
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	workspaceId, err := strconv.ParseInt(workspacesIdParams, 10, 64)
	if workspaceId <= 0 {
		a.notFoundResponse(w, r)
		return
	}
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	var input struct {
		Role string `json:"role"`
	}
	if err = a.readJSON(w, r, &input); err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
	switch input.Role {
	case "admin":
		input.Role = data.AdminPermission
	case "moderator":
		input.Role = data.ModeratorPermission
	case "member":
		input.Role = data.MemberPermission
	}

	if err = a.DB.Permission.UpdateForUser(userId, workspaceId, data.Permissions(input.Role)); err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	if err := a.writeJSON(w, http.StatusOK, envelope{"message": "user role successfully updated"}, nil); err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *app) removeUserFromWorkspaceHandler(w http.ResponseWriter, r *http.Request) {
	workspaceIdParams := chi.URLParam(r, "id")
	userIdParams := chi.URLParam(r, "userId")
	workspaceId, err := strconv.ParseInt(workspaceIdParams, 10, 64)
	if workspaceId <= 0 || err != nil {
		a.notFoundResponse(w, r)
		return
	}
	userId, err := strconv.ParseInt(userIdParams, 10, 64)
	if userId <= 0 || err != nil {
		a.notFoundResponse(w, r)
		return
	}
	if err = a.DB.Permission.RemoveForUser(userId, workspaceId); err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
	if err = a.writeJSON(w, http.StatusOK, envelope{"message": "user successfully removed from workspace"}, nil); err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *app) getAllWorkspaceUsersHandler(w http.ResponseWriter, r *http.Request) {
	params := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(params, 10, 64)
	if id <= 0 {
		a.notFoundResponse(w, r)
		return
	}
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	users, err := a.DB.Workspace.GetAllWorkspaceUsers(id)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
	for _, user := range users {
		switch user.RoleId {
		case 1:
			user.Role = "admin"
		case 2:
			user.Role = "moderator"
		case 3:
			user.Role = "member"
		}
	}

	if err = a.writeJSON(w, http.StatusOK, envelope{"users": users}, nil); err != nil {
		a.serverErrorResponse(w, r, err)
	}
}
