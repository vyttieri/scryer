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

/*
	I was grappling for a bit with a way to do this more efficiently than a query per model.

	gORM doesn't have great support for batch updates. MySQL supports an INSERT INTO ON DUPLICATE KEY
	which will accomplish what I want, but that would require constructing the query string manually.

	gORM also has a .Replace method for associations, but that simply creates new rows and nulls out the association on the old ones.
	That's not ideal either. So I decided to stick with the good old update per row, which isn't the end of the world here,
	as we're only updating 6 rows at a time. If this project were to scale I would defintely want to implemenet a batch solution.
*/
func (user *User) UpdateDevicePreferences(devicePreferences *[]DevicePreference) error {
	var err error

	for _, devicePreference := range *devicePreferences {

		// Using a map for the Updates call is necessary because otherwise gORM will not update "zero" values i.e. 0 or false.
		err := database.Connection.Model(&devicePreference).Where("user_id = ?", user.ID).Updates(map[string]interface{}{
			"icon": devicePreference.Icon,
			"sort_position": devicePreference.SortPosition,
			"visible": devicePreference.Visible,
		}).Error

		if err != nil {
			return err
		}
	}

	return err
}
