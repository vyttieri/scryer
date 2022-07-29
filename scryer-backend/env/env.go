package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var JwtSecretKey, ApiKey, DbHost, DbPort, DbName, DbUsername, DbPassword string

func LoadEnv() {
	fmt.Println("Loading environment")
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	JwtSecretKey = os.Getenv("JWT_SECRET_KEY")
	ApiKey = os.Getenv("ONESTEPGPS_API_KEY")
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbName = os.Getenv("DB_NAME")
	DbUsername = os.Getenv("DB_USERNAME")
	DbPassword = os.Getenv("DB_PASSWORD")

	fmt.Println(ApiKey)
	fmt.Println("Environment loaded")
}
