package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/wonjinsin/go-boilerplate/internal/interfaces"
)

// NewRouter creates and configures a new chi router
func NewRouter(userSvc interfaces.UserService) *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Controllers
	healthCtrl := NewHealthController()
	userCtrl := NewUserController(userSvc)

	// Routes
	r.Get("/healthz", healthCtrl.Check)

	// User routes
	r.Route("/users", func(r chi.Router) {
		r.Post("/", userCtrl.CreateUser)
		r.Get("/", userCtrl.ListUsers)
		r.Get("/{id}", userCtrl.GetUser)
	})

	return r
}
