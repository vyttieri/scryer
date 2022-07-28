package migrations

import	(
	"fmt"

	"gorm.io/gorm"

	client "scryer-backend/db/client"
	"scryer-backend/db/models"
)

var dbInstance *gorm.DB
var dbError error

func Run() {
	dbInstance := client.Connect("scryer:onestepgpsr00lz@tcp(localhost:3306)/scryer")
	fmt.Println("Starting DB Migration")
	dbInstance.AutoMigrate(&models.User{})
	fmt.Println("Finished DB Migration")
}
