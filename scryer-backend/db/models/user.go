package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	database "scryer-backend/db"
)

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	// TODO: Figure out to set length limit with gorm
	Username string	`json:"username" gorm:"index:,unique,class:VARCHAR(24)"`
	Password string `json:"password" gorm:"type:varchar(64)"`

	DevicePreferences []DevicePreference `json:"devicePreferences"`
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

	err := database.Connection.Create(&user).Error

	return err
}

func (user *User) FindByID() error {
	err := database.Connection.Where(&User{Username: user.Username}).First(&user).Error

	return err
}

func (user *User) FindByUsername() error {
	err := database.Connection.Where(&User{Username: user.Username}).First(&user).Error

	return err
}

func (user *User) FindDevicePreferences(devicePreferences *[]DevicePreference) error {
	association := database.Connection.Model(&user).Association("DevicePreferences")
	if err := association.Error; err != nil {
		return err
	}

	// TODO: Find error
	association.Find(&devicePreferences)

	return nil
}

func (user *User) UpdateDevicePreferences(devicePreferences *[]DevicePreference) error {
	return nil
}
