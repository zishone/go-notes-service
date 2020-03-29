package main

import (
	"net/http"

	"github.com/zishone/go-notes-service/cmd/api/routes"
	"github.com/zishone/go-notes-service/internal/platform/configurations"
)

func main() {
	c := configurations.New()

	r := routes.New()
	r.ComposeMiddlewares()
	r.ConfigureRoutes()
	r.WalkRoutes()

	err := http.ListenAndServe(c.Port(), r.Mux())
	if err != nil {
		panic(err)
	}
}
