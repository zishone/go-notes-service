package main

import (
	"fmt"
	"net/http"
	"os"

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
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
