package api

import (
	"fmt"
	"net/http"
)

func Top(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")

	if _, ok := SessionID[name]; !ok {
		fmt.Fprintln(w, "Invalid request")
		return
	}
	renderTemplate(w, "top", name)
}
