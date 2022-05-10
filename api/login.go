package api

import (
	"crypto/sha256"
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	name := r.PostFormValue("name")
	password := r.PostFormValue("password")

	if password == "" || name == "" {
		return
	}

	h := sha256.Sum256([]byte(password))
	hpass := fmt.Sprintf("%x", h)
	fmt.Printf("password is %s\n", password)
	fmt.Printf("hashed password is %s\n", hpass)

	u := User { name, password }

	renderTemplate(w, "top", u)
}
