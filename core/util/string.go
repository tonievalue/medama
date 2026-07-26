package util

import "math/rand/v2"

func GenerateRandomString(length int) string {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	token := make([]byte, length)

	for i := range token {
		b := rand.IntN(len(charset))
		token[i] = charset[b]
	}

	return string(token)
}
