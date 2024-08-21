package main

import (
	"fmt"
	"log"
	"net/http"

	logic "simple-notebook-go/internal/app"
	database "simple-notebook-go/internal/database"
	transport "simple-notebook-go/internal/transport"
)

func main() {

	db := database.ConnectDB()
	defer db.Close()

	database.CreateTable(db)

	notes := logic.GetNote(db)
	for _, note := range notes {
		fmt.Printf("ID: %d, Title: %s, Content: %s\n", note.ID, note.Title, note.Content)
	}

	http.HandleFunc("/notes", transport.NoteHandler(db))
	http.HandleFunc("/getnotes", transport.GetAllNotesHandler(db))
	fmt.Println("Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
