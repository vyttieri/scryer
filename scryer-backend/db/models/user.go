package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	database "scryer-backend/db"
)

// TODO: Add index on Username
type User struct {
	ID uint `json:"id" gorm:"primary_key"`
	Username string	`json:"username" gorm:"unique"`
	Password string `json:"password"`
}

// TODO: Add salt
func (user *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	user.Password = string(bytes)
	return err
}

func (user *User) CheckPassword(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil, err
}

func (user *User) Create() error {
	fmt.Println("calling Create on user")
	fmt.Println(user)
	// TODO: Fix DB stuff

	result := database.Connection.Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}

	return result.Error
}

func (user *User) FindByID() error {
	result := database.Connection.Where(&User{Username: user.Username}).First(&user)
	if result.Error != nil {
		panic(result.Error)
	}

	return result.Error
}

func (user *User) FindByUsername() error {
	result := database.Connection.Where(&User{Username: user.Username}).First(&user)
	if result.Error != nil {
		panic(result.Error)
	}

	return result.Error
}
