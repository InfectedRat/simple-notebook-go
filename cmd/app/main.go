package main

import (
	"fmt"
	"log"
	"net/http"

	logic "simple-notebook-go/internal/app"
	database "simple-notebook-go/internal/database"
	"simple-notebook-go/internal/transport"
)

func main() {

	db := database.ConnectDB()
	defer db.Close()

	database.CreateTable(db)

	notes := logic.GetNote(db)
	for _, note := range notes {
		fmt.Printf("ID: %d, Title: %s, Content: %s\n", note.ID, note.Title, note.Content)
	}

	// http.HandleFunc("/notes", transport.NoteHandler(db))
	// http.HandleFunc("/getnotes", transport.GetAllNotesHandler(db))

	// API маршруты
	http.HandleFunc("/notes", transport.NoteHandler(db))
	http.HandleFunc("/allnotes", transport.GetAllNotesHandler(db))

	// Обслуживаем статические файлы (HTML, CSS, JS)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
