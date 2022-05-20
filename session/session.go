package session

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
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
	n, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		return "", err
	}

	s := fmt.Sprint(n)

	return s, nil
}
