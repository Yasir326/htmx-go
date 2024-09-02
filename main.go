package main

import (
	"fmt"
	"htmx_go/internal/database"
	"htmx_go/internal/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	database.Init()

	router := mux.NewRouter()
	router.HandleFunc("/", handlers.GetFilms).Methods("GET")

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Printf("Server is running on port %s\n", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
