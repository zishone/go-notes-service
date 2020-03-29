package main

import (
	"net/http"

	"github.com/zishone/go-notes-service/cmd/api/routes"
)

// PORT : Port that the service will listen and serve
const PORT = ":3000"

func main() {
	r := routes.New()
	r.ComposeMiddlewares()
	r.ConfigureRoutes()

	err := http.ListenAndServe(PORT, r.Mux())
	if err != nil {
		panic(err)
	}
}
