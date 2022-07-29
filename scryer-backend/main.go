package main

import (
	"flag"

	database "scryer-backend/db"
	"scryer-backend/db/migrations"
	"scryer-backend/env"
	"scryer-backend/server"
)

func main() {
	env.LoadEnv()

	database.Connect()

	flag.Parse()
	cmd := flag.Arg(0)
	switch cmd {
	case "migrate":
		migrations.Run()
	case "server":
		server.Run()
	default:
		server.Run()
	}
}
