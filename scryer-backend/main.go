package main

import (
	"flag"

	"scryer-backend/db"
	"scryer-backend/db/migrations"
	"scryer-backend/env"
	"scryer-backend/server"
)

func main() {
	env.LoadEnv()

	dbContext := db.Connect()

	flag.Parse()
	cmd := flag.Arg(0)
	switch cmd {
	case "migrate":
		migrations.Run(dbContext)
	case "server":
		server.Run(dbContext)
	default:
		server.Run(dbContext)
	}
}
