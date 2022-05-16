package api

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/kunihiro-dev/auth-web-app/model/entity"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Error(w, r)
		return
	}

	name := r.PostFormValue("name")
	password := r.PostFormValue("password")

	if password == "" || name == "" {
		Error(w, r)
		return
	}

	h, err := generatePassword(password, 32)
	if err != nil {
		fmt.Println("Password generate error.")
		Error(w, r)
		return
	}
	hpass := fmt.Sprintf("%x", h)
	fmt.Printf("password is %s\n", password)
	fmt.Printf("hashed password is %s\n", hpass)

	u := entity.UserInfo { Name: name, Password: password }

	n, err  := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		Error(w, r)
		return
	}
	SessionID[u.Name] = n.Int64()
	fmt.Printf("Session ID is %d\n", n)

	Top(w, r)
}

func generatePassword(password string, saltLen int) ([32]byte, error) {
	salt := make([]byte, saltLen)
	var result [32]byte
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return result, err
	}

	passbytes := append(salt, []byte(password)...)
	result = sha256.Sum256(passbytes)
	return result, nil
}
