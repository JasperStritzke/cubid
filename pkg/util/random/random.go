package random

import (
	"math/rand"
)

func GenerateRandom(max int, min int) int {
	return rand.Intn(max-min) + min
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!?.,;:/#")

func GenerateInsecureRandomString(length int) string {
	s := make([]rune, length)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
