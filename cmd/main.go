package main

import (
	"log"
	"net/http"
	"vercel-functions/api"
)

func main() {
	http.HandleFunc("GET /api", api.HttpHandler)

	log.Println("Starting listener... :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
