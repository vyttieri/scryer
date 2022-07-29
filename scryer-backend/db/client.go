package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var Connection *gorm.DB

func Connect(connectionString string) {
	var dbError error
	Connection, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		panic("Something went wrong connecting to database")
	}
}
