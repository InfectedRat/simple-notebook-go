package logic

import (
	"database/sql"
	"fmt"
	"log"

	models "simple-notebook-go/internal/models"
)

func CreateNote(db *sql.DB, title, content string) error {
	query := `INSERT INTO notes (title, content) VALUES (?, ?)`
	_, err := db.Exec(query, title, content)
	if err != nil {
		fmt.Errorf("Запрос завершился ошибкой: %v", err)
	}
	fmt.Println("Заметка успешно создана!")
	return nil
}

func GetNote(db *sql.DB) []models.Note {
	rows, err := db.Query("SELECT id, title, content FROM notes")
	if err != nil {
		log.Fatalf("Ошибка выполнения запроса: %v", err)
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var note models.Note
		err := rows.Scan(&note.ID, &note.Title, &note.Content)
		if err != nil {
			log.Fatalf("Ошибка при чтении заметки: %v", err)
		}
		notes = append(notes, note)
	}
	return notes
}
