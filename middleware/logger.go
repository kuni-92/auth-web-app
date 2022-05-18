package middleware

import (
	"log"
	"net/http"
)

func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f := "[%s] %s\n"
		log.Printf(f, r.Method, r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}
