package services

import (
	"regexp"
	"net/mail"
	"golang.org/x/crypto/bcrypt"
)

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func containsSQLInjection(input string) bool {
	re := regexp.MustCompile(`(?i)(\b(SELECT|INSERT|UPDATE|DELETE|DROP|UNION|--|;)\b)`)
	return re.MatchString(input)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}