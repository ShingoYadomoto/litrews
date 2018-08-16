package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func ToHash(p string) (h string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	h = string(hash)

	return
}

func CompareHashAndPlain(h string, p string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(h), []byte(p))

	return
}
