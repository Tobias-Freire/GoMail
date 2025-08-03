package main

import (
	"os"
	"log"
	"net/http"

	"gomail/config"
	"gomail/handlers"
)

func main() {
	config.LoadEnv()
	port := os.Getenv("API_PORT")

	http.HandleFunc("/send-email", handlers.EmailHandler)

	log.Println("Server listening and serving on "+port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}