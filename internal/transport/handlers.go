package transport

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	logic "simple-notebook-go/internal/app"
)

type NoteJSON struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NoteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
			return
		}

		var note NoteJSON
		err := json.NewDecoder(r.Body).Decode(&note)
		if err != nil {
			http.Error(w, "Невозможно обработать запрос", http.StatusBadRequest)
			return
		}

		err = logic.CreateNote(db, note.Title, note.Content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Заметка успешно создана")
	}
}

func GetAllNotesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
			return
		}

		// Получаем все заметки, используя вашу функцию GetNote
		notes := logic.GetNote(db)

		// Кодируем заметки в JSON и отправляем ответ
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(notes)
		if err != nil {
			http.Error(w, "Ошибка при формировании ответа", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
