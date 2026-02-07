package api

import (
	"fmt"
	"net/http"
	"vercel-functions/api/.vercel/db"
)

func HttpDateHandler(w http.ResponseWriter, r *http.Request) {

	DB := db.NewDB()

	response := fmt.Sprintf("Hello from the DB: %s", DB.Name)

	w.Write([]byte(response))
}
