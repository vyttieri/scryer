package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


	var dbInstance *gorm.DB
	var dbError error

func Connect(connectionString string) (dbInstance *gorm.DB) {
	dbInstance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		panic("Something went wrong connecting to database")
	}

	return dbInstance
}
