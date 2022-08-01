package migrations

import	(
	"fmt"

	database "scryer-backend/db"
	"scryer-backend/db/models"
)

func Run() {
	fmt.Println("Starting DB Migration")
	database.Connection.AutoMigrate(&models.User{})
	database.Connection.AutoMigrate(&models.DevicePreference{})
	fmt.Println("Finished DB Migration")
}
