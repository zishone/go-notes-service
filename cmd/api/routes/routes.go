package routes

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Router : A struct that will hold the mux
type Router struct {
	mux *chi.Mux
}

// New : Instantiates the router
func New() *Router {
	return &Router{
		mux: chi.NewRouter(),
	}
}

// Mux : Getter for mux
func (r *Router) Mux() *chi.Mux {
	return r.mux
}

// ComposeMiddlewares : Composes global middlewares
func (r *Router) ComposeMiddlewares() {
	r.mux.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
	)
}

// ConfigureRoutes : Configure api routes
func (r *Router) ConfigureRoutes() {
	r.mux.Route("/api", func(r chi.Router) {
		r.Route("/v1", v1)
	})
}

// WalkRoutes : Walk and print out all routes
func (r *Router) WalkRoutes() error {
	err := chi.Walk(r.mux, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// Init : Initializes the router
func (r *Router) Init() error {
	r.ComposeMiddlewares()
	r.ConfigureRoutes()
	err := r.WalkRoutes()
	if err != nil {
		return err
	}
	return nil
}
