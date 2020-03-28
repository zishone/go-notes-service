package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/zishone/go-notes-service/cmd/api/routes"
)

func main() {
	r := routes.Init()
	port := ":3000"
	err := http.ListenAndServe(port, r)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
