package api

import "net/http"

func HttpDateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Date!"))
}
