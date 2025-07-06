package todo

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func Todos(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			rows, err := db.Query("SELECT id, text, done FROM todos")
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			defer rows.Close()

			var todos []Todo
			for rows.Next() {
				var t Todo
				if err := rows.Scan(&t.ID, &t.Text, &t.Done); err != nil {
					http.Error(w, err.Error(), 500)
					return
				}
				todos = append(todos, t)
			}
			json.NewEncoder(w).Encode(todos)

		case http.MethodPost:
			var t Todo
			if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
			res, err := db.Exec("INSERT INTO todos (text, done) VALUES (?, ?)", t.Text, t.Done)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			id, _ := res.LastInsertId()
			t.ID = int(id)
			json.NewEncoder(w).Encode(t)

		default:
			http.Error(w, "Method not allowed", 405)
		}
	}
}

func TodoByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", 400)
			return
		}

		switch r.Method {
		case http.MethodGet:
			var t Todo
			err := db.QueryRow("SELECT id, text, done FROM todos WHERE id = ?", id).Scan(&t.ID, &t.Text, &t.Done)
			if err == sql.ErrNoRows {
				http.NotFound(w, r)
				return
			} else if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			json.NewEncoder(w).Encode(t)

		case http.MethodPut:
			var t Todo
			if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
			_, err := db.Exec("UPDATE todos SET text = ?, done = ? WHERE id = ?", t.Text, t.Done, id)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			t.ID = id
			json.NewEncoder(w).Encode(t)

		case http.MethodDelete:
			_, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.WriteHeader(http.StatusNoContent)

		default:
			http.Error(w, "Method not allowed", 405)
		}
	}
}
