package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"scryer-backend/env"
)


var Connection *gorm.DB

func Connect() {
	fmt.Println("hello")
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		env.DbUser,
		env.DbUserPassword,
		env.DbHost,
		env.DbPort,
		env.DbName,
	)
	fmt.Println("Connecting to database: ", connectionString)

	var dbError error
	Connection, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		panic("Something went wrong connecting to database")
	}
}
