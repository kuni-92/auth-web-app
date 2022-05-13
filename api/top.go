package api

import (
	"net/http"
)

func Top(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")

	if _, ok := SessionID[name]; !ok {
		Error(w ,r)
		return
	}
	renderTemplate(w, "top", name)
}
