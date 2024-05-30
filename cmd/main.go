package main

import (
	"log"
	"net/http"
	"os"

	"github.com/abhi-nolyfe/chess-goblins/internal/handlers"
	"github.com/joho/godotenv"
)

func main () {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: mux,
	}

	mux.HandleFunc("/", handlers.HandleHome)

	if portString == "" {
		log.Fatal("PORT is not found in the environment.")
	}
	log.Printf("Serving on port: %s\n", portString)
	log.Fatal(srv.ListenAndServe())
}