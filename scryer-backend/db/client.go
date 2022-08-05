package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"scryer-backend/env"
)

type DbContext struct {
	connectionString string
	Connection *gorm.DB
}
func Connect() *DbContext {
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		env.DbUser,
		env.DbUserPassword,
		env.DbHost,
		env.DbPort,
		env.DbName,
	)
	fmt.Println("Connecting to database: ", connectionString)

	Connection, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		fmt.Println("Couldnt connect to db", err)
		panic(err)
	}

	return &DbContext{connectionString: connectionString, Connection: Connection}
}
