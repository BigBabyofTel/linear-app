package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/lucabrx/wuhu/internal/data"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *app) routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(httprate.LimitByIP(100, 1*time.Minute))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Use(a.authenticate)

	r.NotFound(a.notFoundResponse)
	r.MethodNotAllowed(a.methodNotAllowedResponse)

	r.Route("/v1/auth", func(r chi.Router) {
		r.Post("/register", a.emailRegisterUserHandler)
		r.Post("/login", a.emailLoginUserHandler)
		r.Post("/logout", a.requireAuthenticatedUser(a.logoutHandler))
	})

	r.Route("/v1/user", func(r chi.Router) {
		r.Get("/", a.requireAuthenticatedUser(a.getSessionUserHandler))
		r.Delete("/", a.requireAuthenticatedUser(a.deleteUserAccountHandler))
		r.Patch("/password", a.requireAuthenticatedUser(a.updatePasswordHandler))
	})

	r.Route("/v1/issues", func(r chi.Router) {
		r.Post("/", a.requireAuthenticatedUser(a.createIssueHandler))
	})

	r.Route("/v1/upload", func(r chi.Router) {
		r.Post("/image", a.requireAuthenticatedUser(a.uploadImageHandler))
	})

	r.Route("/v1/workspaces", func(r chi.Router) {
		r.Post("/", a.requireAuthenticatedUser(a.createWorkspaceHandler))
		r.Get("/{slug}", a.requirePermission(data.MemberPermission, a.requireAuthenticatedUser(a.getWorkspaceHandler)))
		r.Delete("/{slug}", a.requirePermission(data.AdminPermission, a.requireAuthenticatedUser(a.deleteWorkspaceHandler)))
		r.Patch("/{slug}", a.requirePermission(data.AdminPermission, a.requireAuthenticatedUser(a.updateWorkspaceHandler)))
		r.Get("/users/{slug}/{id}", a.requirePermission(data.MemberPermission, a.requireAuthenticatedUser(a.getAllWorkspaceUsersHandler)))
		r.Delete("/remove-users/{id}/{userId}", a.requirePermission(data.AdminPermission, a.requireAuthenticatedUser(a.removeUserFromWorkspaceHandler)))
		r.Patch("/update-role/{slug}/{id}/{userId}", a.requirePermission(data.AdminPermission, a.requireAuthenticatedUser(a.changeUserRoleHandler)))
		r.Get("/my-workspaces", a.requireAuthenticatedUser(a.getAllUserWorkspacesHandler))
		r.Delete("/delete-image/{slug}", a.requirePermission(data.AdminPermission, a.requireAuthenticatedUser(a.removeWorkspaceImageHandler)))
	})

	r.Route("/v1/requests", func(r chi.Router) {
		r.Post("/invite/{slug}", a.requirePermission(data.AdminPermission, a.requireAuthenticatedUser(a.inviteUserToWorkspaceHandler)))
		r.Get("/", a.requireAuthenticatedUser(a.getAllRequestsHandler))
		r.Post("/", a.requireAuthenticatedUser(a.answerRequestHandler))
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusMovedPermanently)
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("doc.json"),
	))

	return r
}
