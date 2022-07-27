package main

import "scryer-backend/db"

func main() {
	db.Connect("scryer:onestepgpsr00lz@tcp(localhost:3306)/scryer")
	db.Migrate()
}
