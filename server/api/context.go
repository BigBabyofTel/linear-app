package api

import (
	"context"
	"net/http"

	"github.com/lucabrx/wuhu/internal/data"
)

type contextKey string

const userContextKey = contextKey("user")
const tokenContextKey = contextKey("token")
const workspaceContextKey = contextKey("workspace")

func (a *app) contextSetUser(r *http.Request, user *data.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

func (a *app) contextGetUser(r *http.Request) *data.User {
	user, ok := r.Context().Value(userContextKey).(*data.User)

	if !ok {
		panic("missing user value in request context")
	}
	return user
}

func (a *app) contextSetToken(r *http.Request, token string) *http.Request {
	ctx := context.WithValue(r.Context(), tokenContextKey, token)
	return r.WithContext(ctx)
}

func (a *app) contextGetToken(r *http.Request) string {
	token, ok := r.Context().Value(tokenContextKey).(string)

	if !ok {
		panic("missing token value in request context")
	}
	return token
}

func (a *app) contextSetWorkspace(r *http.Request, workspace *data.Workspace) *http.Request {
	ctx := context.WithValue(r.Context(), workspaceContextKey, workspace)
	return r.WithContext(ctx)
}

func (a *app) contextGetWorkspace(r *http.Request) *data.Workspace {
	workspace, ok := r.Context().Value(workspaceContextKey).(*data.Workspace)

	if !ok {
		panic("missing workspace value in request context")
	}
	return workspace
}
