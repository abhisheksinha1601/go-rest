package common

import (
	"crypto/sha512"
	"encoding/base64"

	"golang.org/x/crypto/pbkdf2"
)

const iterations = 10000

func Pbkdf2Encrypt(password, salt string) string {
	return base64.StdEncoding.EncodeToString([]byte(pbkdf2.Key([]byte(password), []byte(salt), iterations, 64, sha512.New)))
}
