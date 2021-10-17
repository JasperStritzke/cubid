package util

import (
	"github.com/jasperstritzke/cubid/pkg/util/random"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password, salt string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password+salt), 14)
	return string(bytes)
}

func CheckPasswordHash(password, salt, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+salt))
	return err == nil
}

func HashString(input string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(input), 14)

	if err != nil {
		return ""
	}

	return string(bytes)
}

const saltLength = 32

func Salt() string {
	salt, _ := random.GenerateRandomString(saltLength)
	return salt
}
