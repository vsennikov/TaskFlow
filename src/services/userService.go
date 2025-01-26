package services

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/vsennikov/TaskFlow/src/repository"
)

type UserServiceInterface interface {
	UserRegistration(name string, email string, password string) (uint, error)
	Login(email string, password string) (string, error)
	DecodeToken(tokenString string) (uint, error)
}

type UserService struct {
	repository repository.UserDBInterface
}

func NewUserService(r repository.UserDBInterface) *UserService {
	return &UserService{
		repository: r,
	}
}

func (u *UserService) UserRegistration(name string, email string, password string) (uint, error) {
	if name == "" || email == "" || password == "" {
		return 0, errors.New("name, email, and password cannot be empty")
	}
	if !isEmailValid(email) {
		return 0, errors.New("invalid email")
	}
	if containsSQLInjection(name) || containsSQLInjection(email) || containsSQLInjection(password) {
		return 0, errors.New("input contains invalid characters")
	}
	_, err := u.repository.GetUserByEmail(email)
	if err == nil {
		return 0, errors.New("email already exists")
	}
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return 0, err
	}
	return u.repository.CreateUser(name, email, hashedPassword)
}

func (u *UserService) Login(email string, password string) (string, error) {
	if email == "" || password == "" {
		return "", errors.New("email and password cannot be empty")
	}
	if !isEmailValid(email) {
		return "", errors.New("invalid email")
	}
	if containsSQLInjection(email) || containsSQLInjection(password) {
		return "", errors.New("input contains invalid characters")
	}
	user, err := u.repository.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}
	if !checkPasswordHash(password, user.Password) {
		return "", errors.New("invalid password")
	}
	token, err := generateJWT(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func generateJWT(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (u *UserService) DecodeToken(tokenString string) (uint, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["user_id"].(float64))
		return userID, nil
	}
	return 0, errors.New("invalid token")
}
