package utils

import (
	"crypto/rand"
	"math/big"
)

const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		b[i] = charset[n.Int64()]
	}
	return string(b)
}
