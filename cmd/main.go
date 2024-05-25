package main

import (
	"log"
	"net/http"
	"os"

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

	if portString == "" {
		log.Fatal("PORT is not found in the environment.")
	}
	log.Printf("Serving on port: %s\n", portString)
	log.Fatal(srv.ListenAndServe())
}