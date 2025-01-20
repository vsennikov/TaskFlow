package repository

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserDBInterface interface {
	CreateUser(name string, email string, password string) (uint, error)
	GetUserByEmail(email string) (user, error)
}

type UserDB struct {
}

type user struct {
	gorm.Model
	Username string `json:"username" gorm:"column:username"`
	Email    string `json:"email" gorm:"column:email;unique"`
	Password string `json:"password" gorm:"column:password"`
}

func (user) TableName() string {
	return "users"
}

var postgresqlURL string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	postgresqlURL = "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASS") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=" + os.Getenv("DB_SSLMODE")
}

func getUserDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(postgresqlURL))
	if err != nil {
		panic("failed to connect database")
	}
	_ = db.AutoMigrate(&user{})

	return db
}

func (*UserDB) CreateUser(name string, email string, password string) (uint, error) {
	user := user{Username: name, Email: email, Password: password}
	err := getUserDB().Save(&user).Error

	return user.Model.ID, err
}

func (*UserDB) GetUserByEmail(email string) (user, error) {
	var user user
	err := getUserDB().Where("email = ?", email).First(&user).Error
	return user, err
}