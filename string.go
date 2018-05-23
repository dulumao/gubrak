package gubrak

import (
	randCrypto "crypto/rand"
	"encoding/base64"
	randMath "math/rand"
	"time"
)

func init() {
	randMath.Seed(time.Now().UnixNano())
}

// RandomString function generate random string in defined length
func RandomString(length int) string {
	data := make([]byte, length)

	if _, err := randCrypto.Read(data); err != nil {
		return ""
	}

	res := base64.URLEncoding.EncodeToString(data)
	return res
}

// RandomStringAlphabet function generate random alphabet string in defined length
func RandomStringAlphabet(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randMath.Intn(len(letters))]
	}

	return string(b)
}
