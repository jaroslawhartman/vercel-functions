package api

import (
	"fmt"
	"net/http"
	"strings"
)

func HttpEchoHandler(w http.ResponseWriter, r *http.Request) {

	ix := 0
	output := ""

	input := r.URL.Query().Get("input")

	for _, ch := range input {
		r := string(ch)
		if ix%2 == 0 {
			r = strings.ToUpper(r)
		} else {
			r = strings.ToLower(r)
		}

		ix++
		output = output + string(r)

		fmt.Println("> " + output)
	}

	fmt.Println(output)

	response := fmt.Sprintf("{ \"input\": \"%s\", \"output\": \"%s\"}", input, output)

	w.Write([]byte(response))
}
