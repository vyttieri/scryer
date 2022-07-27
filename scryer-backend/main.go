package main

import (
	"scryer-backend/db"
	"scryer-backend/server"
)

func main() {
	switch "" {
	case "migrate":
		db.RunMigrations()
	case "server":
		server.Run()
	default:
		server.Run()
	}
}
