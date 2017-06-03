package csrf

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
)

const (
	chars   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	idxBits = 6
	idxMask = 1<<idxBits - 1
)

// GenerateRandom generates a random string of specified length.
func GenerateRandom(length int) string {
	result := make([]byte, length)
	bufferSize := int(float64(length) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < length; j++ {
		if j%bufferSize == 0 {
			randomBytes = secureRandomBytes(bufferSize)
		}
		if idx := int(randomBytes[j%length] & idxMask); idx < len(chars) {
			result[i] = chars[idx]
			i++
		}
	}

	return string(result)
}

// GenerateToken generates a secure token from a secret and salt.
func GenerateToken(secret, salt string) string {
	return salt + hash(salt+"-"+secret)
}

// Verify verifies if a token is valid.
// It takes in the salt length and secret used to create it.
func Verify(token, secret string, saltLen int) bool {
	salt := token[0:saltLen]
	return salt+hash(salt+"-"+secret) == token
}

// hash hashes a string using sha256 and returns a
// base64 encoded string, which is URL safe.
func hash(str string) string {
	hash := sha256.New()
	io.WriteString(hash, str)

	hashedString := base64.RawURLEncoding.EncodeToString(hash.Sum(nil))
	return hashedString
}

// secureRandomBytes generates secure random bytes of a specified
// length.
func secureRandomBytes(length int) []byte {
	var randomBytes = make([]byte, length)
	rand.Read(randomBytes)
	return randomBytes
}
