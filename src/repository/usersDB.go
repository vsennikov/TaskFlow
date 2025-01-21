package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserDBInterface interface {
	CreateUser(name string, email string, password string) (uint, error)
	GetUserByEmail(email string) (user, error)
}

type UserDB struct {
}

func (user) TableName() string {
	return "users"
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
