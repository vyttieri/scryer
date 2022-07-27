package main

import (
	"flag"

	"scryer-backend/db/migrations"
	"scryer-backend/server"
)

func main() {
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
