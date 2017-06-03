// Package csrf is a logic package for creating CSRF middleware.
package csrf

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
)

type Toolset struct {
	secret  string
	saltLen int
}

type Options struct {
	Secret  string
	SaltLen int
}

const (
	chars   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	idxBits = 6
	idxMask = 1<<idxBits - 1
)

// GenerateSalt generates a random string of specified length.
func (f *Toolset) GenerateSalt() string {
	result := make([]byte, f.saltLen)
	bufferSize := int(float64(f.saltLen) * 1.3)
	for i, j, randomBytes := 0, 0, []byte{}; i < f.saltLen; j++ {
		if j%bufferSize == 0 {
			randomBytes = secureRandomBytes(bufferSize)
		}
		if idx := int(randomBytes[j%f.saltLen] & idxMask); idx < len(chars) {
			result[i] = chars[idx]
			i++
		}
	}

	return string(result)
}

// GenerateToken generates a secure token from a secret and salt.
func (f *Toolset) GenerateToken(salt string) string {
	return salt + hash(salt+"-"+f.secret)
}

// Verify verifies if a token is valid.
// It takes in the salt length and secret used to create the token.
func (f *Toolset) Verify(token string) bool {
	salt := token[0:f.saltLen]
	return salt+hash(salt+"-"+f.secret) == token
}

// New returns a new Toolset, it takes in a type Options.
// The toolset will use the options.
func New(opt Options) *Toolset {
	return &Toolset{
		secret:  opt.Secret,
		saltLen: opt.SaltLen,
	}
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
