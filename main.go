package main

import (
	"log"
	"net/http"

	"github.com/its-asif/go-url-shortener/config"
	"github.com/its-asif/go-url-shortener/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	config.InitDB()

	r := mux.NewRouter()

	r.HandleFunc("/shorten", handlers.ShortenURL).Methods("POST")
	r.HandleFunc("/{code}", handlers.RedirectURL).Methods("GET")

	log.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
