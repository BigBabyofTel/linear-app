package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"

	"github.com/lucabrx/wuhu/internal/data"
	"github.com/lucabrx/wuhu/internal/validator"
)

func (a *app) requireAuthenticatedUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := a.contextGetUser(r)
		if user.IsAnonymous() {
			a.authenticationRequiredResponse(w, r)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func (a *app) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" || !strings.HasPrefix(tokenHeader, "Bearer ") {
			r = a.contextSetUser(r, data.AnonymousUser)
			next.ServeHTTP(w, r)
			return
		}
		token := strings.TrimPrefix(tokenHeader, "Bearer ")
		r = a.contextSetToken(r, token)
		v := validator.New()
		if data.ValidateTokenPlaintext(v, token); !v.Valid() {
			a.invalidAuthenticationTokenResponse(w, r)
			return
		}

		user, err := a.DB.User.GetForToken(data.ScopeAuthentication, token)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrRecordNotFound):
				a.invalidAuthenticationTokenResponse(w, r)
			default:
				a.serverErrorResponse(w, r, err)
			}
			return
		}

		r = a.contextSetUser(r, user)
		next.ServeHTTP(w, r)
	})
}

func (a *app) requirePermission(code string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := a.contextGetUser(r)
		slug := chi.URLParam(r, "slug")
		params := chi.URLParam(r, "id")
		id, _ := strconv.ParseInt(params, 10, 64)

		workspace, err := a.DB.Workspace.Get(id, slug, session.ID)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrRecordNotFound):
				a.notFoundResponse(w, r)
			default:
				a.serverErrorResponse(w, r, err)
			}
			return
		}

		permission, err := a.DB.Permission.GetForUser(session.ID, workspace.ID)
		if err != nil {
			a.serverErrorResponse(w, r, err)
			return
		}

		if !permission.Contains(code) {
			a.notPermittedResponse(w, r)
			return
		}

		r = a.contextSetWorkspace(r, workspace)
		next.ServeHTTP(w, r)
	}
}
