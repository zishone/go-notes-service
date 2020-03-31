package routes

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type router struct {
	mux *chi.Mux
}

func (r *router) configureRoutes() {
	r.mux.Route("/api", func(r chi.Router) {
		r.Route("/v1", rv1)
	})
}

func (r *router) composeMiddlewares() {
	r.mux.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
	)
}

func (r *router) walkRoutes() error {
	err := chi.Walk(r.mux, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// NewRouter : Instantiates a new router
func NewRouter() (*chi.Mux, error) {
	r := router{mux: chi.NewRouter()}
	r.composeMiddlewares()
	r.configureRoutes()
	if err := r.walkRoutes(); err != nil {
		return nil, err
	}
	return r.mux, nil
}
