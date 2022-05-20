package api

import "net/http"

func Error(w http.ResponseWriter, r *http.Request) {
	if c, err := r.Cookie(SESSION_KEY); err == nil {
		c.MaxAge = -1
		http.SetCookie(w, c)
	}
	renderTemplate(w, "error", nil)
}
