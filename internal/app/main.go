package logic

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateNote(db *sql.DB, title, content string) {
	query := `INSERT INTO notes (title, content) VALUES (?, ?)`
	_, err := db.Exec(query, title, content)
	if err != nil {
		log.Fatalf("Запрос завершился ошибкой: %v", err)
	}
	fmt.Println("Заметка успешно создана!")
}
