package router

import (
	"database/sql"
	"golang-rest/internals/model/todo"
	"net/http"
)

func New(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/todos", todo.Todos(db))
	mux.HandleFunc("/todos/", todo.TodoByID(db))

	return mux
}
