package session

import "errors"

var session map[string]int64

func init() {
	session = make(map[string]int64)
}

func SetSession(key string, value int64) error {
	// return an error, if session data exists
	if _, ok := session[key]; ok {
		return errors.New("Session data is exists.")
	}
	session[key] = value
	return nil
}

func GetSession(key string) (int64, error) {
	// return an error, if session data not exists
	if v, ok := session[key]; ok {
		return v, nil
	}
	return 0, errors.New("Session data is not exists.")
}
