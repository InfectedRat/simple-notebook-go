package main

import (
	"fmt"
	logic "simple-notebook-go/internal/app"
	database "simple-notebook-go/internal/database"
)

func main() {

	db := database.ConnectDB()
	defer db.Close()

	database.CreateTable(db)

	logic.CreateNote(db, "Вторая заметка", "Содержимое второй заметки")

	notes := logic.GetNote(db)
	for _, note := range notes {
		fmt.Printf("ID: %d, Title: %s, Content: %s\n", note.ID, note.Title, note.Content)
	}

}
