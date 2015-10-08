package core

import (
	"crypto/rand"
	"encoding/base64"
	bc "golang.org/x/crypto/bcrypt"
)

func GenerateBytes(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GenerateStrings(size int) (string, error) {
	b, err := GenerateBytes(size)
	return base64.URLEncoding.EncodeToString(b), err

}

func Encrypt(p *string) (string, error) {
	salt, _ := GenerateStrings(32)
	ps := []byte(*p + salt)
	hash, e := bc.GenerateFromPassword(ps, bc.DefaultCost)
	*p = string(hash)
	return salt, e
}

func CompareHash(p string, s string, h string) bool {
	err := bc.CompareHashAndPassword([]byte(h), []byte(p+s))
	return err == nil
}
