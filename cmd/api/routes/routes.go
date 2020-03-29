package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/zishone/go-notes-service/cmd/api/handlers"
)

// Router : A struct that will hold the mux
type Router struct {
	mux *chi.Mux
}

// New : Instantiates the router
func New() Router {
	return Router{
		mux: chi.NewRouter(),
	}
}

// Mux : Getter for mux
func (r Router) Mux() *chi.Mux {
	return r.mux
}

// ComposeMiddlewares : Composes global middlewares
func (r Router) ComposeMiddlewares() {
	r.mux.Use(middleware.RequestID)
	r.mux.Use(middleware.RealIP)
	r.mux.Use(middleware.Logger)
	r.mux.Use(middleware.Recoverer)
}

// ConfigureRoutes : Configure api routes
func (r Router) ConfigureRoutes() {
	r.mux.Get("/notes", handlers.GetNotes)
}
