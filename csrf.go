package csrf

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"io"
)

func GenerateUid(length int) string {
	token := make([]byte, length)
	rand.Read(token)

	tokenString := base64.RawURLEncoding.EncodeToString(token)
	return tokenString
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
