package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
)

/* EXAMPLE
func main() {
	secret := GenerateRandom(32)

	for i := 0; i < 80; i++ {
		salt := GenerateRandom(16)
		GenerateToken(secret, salt)
	}
}*/

func GenerateRandom(length int) string {
	random := make([]byte, length)
	rand.Read(random)

	randomString := base64.RawURLEncoding.EncodeToString(random)
	return randomString
}

func GenerateToken(secret, salt string) string {
	return salt + "-" + hash(salt+"-"+secret)
}

func hash(str string) string {
	hash := sha256.New()
	io.WriteString(hash, str)

	hashedString := base64.RawURLEncoding.EncodeToString(hash.Sum(nil))
	return hashedString
}
