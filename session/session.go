package session

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

var session map[string]string

func init() {
	session = make(map[string]string)
}

func Add(key string, value string) error {
	// return an error, if session data exists
	if _, ok := session[key]; ok {
		return errors.New("Session data is exists.")
	}
	session[key] = value
	fmt.Printf("add session \n %v\n", session)
	return nil
}

func Take(key string) (string, error) {
	fmt.Printf("key is %s take session \n %v\n", key, session)
	// return an error, if session data not exists
	if v, ok := session[key]; ok {
		return v, nil
	}
	return "", errors.New("Session data is not exists.")
}


func Generate() (string, error) {
	s := make([]byte, 128)
	_, err := io.ReadFull(rand.Reader, s)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(s), nil
}
