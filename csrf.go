package csrf

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

func GenerateRandom(len int) string {
	random := make([]byte, len)
	rand.Read(random)

	randomString := base64.RawURLEncoding.EncodeToString(random)
	return randomString
}

func GenerateToken(secret, salt string) string {
	return salt + hash(salt+"-"+secret)
}

func hash(str string) string {
	hash := sha256.New()
	io.WriteString(hash, str)

	hashedString := base64.RawURLEncoding.EncodeToString(hash.Sum(nil))
	return hashedString
}

func Verify(token, secret string, saltLen int) bool {
	salt := token[0:saltLen]
	return salt+hash(salt+"-"+secret) == token
}
