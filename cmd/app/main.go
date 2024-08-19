package main

import (
	logic "simple-notebook-go/internal/app"
	database "simple-notebook-go/internal/database"
)

func main() {

	db := database.ConnectDB()
	defer db.Close()

	database.CreateTable(db)

	logic.CreateNote(db, "Первая заметка", "Содержимое первой заметки")

}
