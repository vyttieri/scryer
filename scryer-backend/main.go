package main

import (
	"flag"

	"scryer-backend/db"
	"scryer-backend/server"
)

func main() {
	flag.Parse()
	cmd := flag.Arg(0)
	switch cmd {
	case "migrate":
		db.RunMigrations()
	case "server":
		server.Run()
	default:
		server.Run()
	}
}
