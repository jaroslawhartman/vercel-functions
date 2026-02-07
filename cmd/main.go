package main

import (
	"log"
	"net/http"
	"vercel-functions/api"
)

func main() {
	http.HandleFunc("GET /api/hello", api.HttpHelloHandler)
	http.HandleFunc("GET /api/date", api.HttpDateHandler)

	log.Println("Starting listener... :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
