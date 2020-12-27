package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lukejmann/detach2-backend/api-server/api"
)

const portStr = ":80"

func main() {
	m := http.NewServeMux()

	apiHandler := api.Handler()
	m.Handle("/", apiHandler)

	fmt.Printf("API Server Listening on port %v\n", portStr)

	err := http.ListenAndServe(portStr, m)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
