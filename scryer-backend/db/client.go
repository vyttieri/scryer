package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"scryer-backend/db/models"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) () {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		panic("Something went wrong connecting to database")
	}
}

func Migrate() {
	fmt.Println("Starting DB Migration")
	Instance.AutoMigrate(&models.User{})
	fmt.Println("Finished DB Migration")
}

func RunMigrations() {
	Connect("scryer:onestepgpsr00lz@tcp(localhost:3306)/scryer")
	Migrate()
}