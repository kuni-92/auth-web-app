package main

import (
	"log"
	"net/http"

	"github.com/kunihiro-dev/auth-web-app/api"
)

func main() {
	http.HandleFunc("/", api.Index)

	if err := http.ListenAndServe(":3030", nil); err != nil {
		log.Println("http server error.")
	}
}
