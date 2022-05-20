package session

import (
	"errors"
	"crypto/rand"
	"math/big"

)

var session map[string]int64

func init() {
	session = make(map[string]int64)
}

func Add(key string, value int64) error {
	// return an error, if session data exists
	if _, ok := session[key]; ok {
		return errors.New("Session data is exists.")
	}
	session[key] = value
	return nil
}

func Take(key string) (int64, error) {
	// return an error, if session data not exists
	if v, ok := session[key]; ok {
		return v, nil
	}
	return 0, errors.New("Session data is not exists.")
}


func Generate() (int64, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		return 0, err
	}
	return n.Int64(), nil
}
