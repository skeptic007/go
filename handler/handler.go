package handler

import (
	"fmt"
	"html"
	"net/http"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q %q", html.EscapeString(r.URL.Path), r.URL.Query())
}
