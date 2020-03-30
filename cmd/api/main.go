package main

import (
	"net/http"

	"github.com/zishone/go-notes-service/cmd/api/routes"
	"github.com/zishone/go-notes-service/internal/platform/configurations"
)

func main() {
	r := routes.New()

	if err := r.Init(); err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(configurations.Port(), r.Mux()); err != nil {
		panic(err)
	}
}
