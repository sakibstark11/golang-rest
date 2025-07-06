package main

import (
	"golang-rest/internals/database"
	"golang-rest/internals/router"
	"log"
	"net/http"
)

func main() {
	database, err := database.Init()
	if err != nil {
		log.Fatal("failed to connect to DB:", err)
	}
	defer database.Close()

	mux := router.New(database)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
