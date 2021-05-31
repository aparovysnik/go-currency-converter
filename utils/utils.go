package utils

import (
	"crypto/rand"
	"math/big"

	"crypto/sha256"
)

//ComputeHash calculates hash of a string
func ComputeHash(input string) []byte {
	hash := sha256.Sum256([]byte(input))

	return hash[:]
}

//GenerateRandomString returns a random string of n chars
func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

//ContainsString checks if a string slice contains a string
func ContainsString(collection []string, value string) bool {
	for _, v := range collection {
		if v == value {
			return true
		}
	}

	return false
}
