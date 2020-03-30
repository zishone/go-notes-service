package main

import (
	"net/http"

	"github.com/zishone/go-notes-service/cmd/api/routes"
	"github.com/zishone/go-notes-service/internal/platform/configurations"
)

func main() {
	r, err := routes.New()
	if err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(configurations.Port(), r); err != nil {
		panic(err)
	}
}
