package main

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"io"
)

/* EXAMPLE
func main() {
	const secret = "BpWgNVso4QWV4RDxSpuKqTOzIzrvbpqLz2Laj3ivCIE"

	for i := 0; i < 80; i++ {
		salt := GenerateRandomString(8)
		fmt.Println(GenerateToken(secret, salt))
	}
}
*/

func GenerateRandomString(length int) string {
	random := make([]byte, length)
	rand.Read(random)

	randomString := base64.RawURLEncoding.EncodeToString(random)
	return randomString
}

func GenerateToken(secret, salt string) string {
	return salt + "-" + hash(salt+"-"+secret)
}

func hash(str string) string {
	hash := sha1.New()
	io.WriteString(hash, str)

	hashedString := base64.RawURLEncoding.EncodeToString(hash.Sum(nil))
	return hashedString
}
