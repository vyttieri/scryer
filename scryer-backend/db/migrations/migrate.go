package migrations

import	(
	"fmt"

	"scryer-backend/db"
	"scryer-backend/db/models"
)

func Run(dbContext *db.DbContext) {
	fmt.Println("Starting DB Migration")
	dbContext.Connection.AutoMigrate(&models.User{})
	dbContext.Connection.AutoMigrate(&models.DevicePreference{})
	fmt.Println("Finished DB Migration")
}
