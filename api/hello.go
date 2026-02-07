package api

import "net/http"

func HttpHelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}
