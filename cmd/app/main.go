package main

import (
	database "simple-notebook-go/internal/database"
)

func main() {

	db := database.ConnectDB()
	defer db.Close()

	database.CreateTable(db)

}
