package router

import (
	"database/sql"
	"golang-rest/internals/model/todo"
	"net/http"
)

func New(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	todoHandler := todo.NewHandler(db)

	mux.HandleFunc("/todos", todoHandler.Todos)
	mux.HandleFunc("/todos/{id}", todoHandler.TodoByID)

	return mux
}
