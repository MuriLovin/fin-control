package main

import (
	"log"
	"net/http"
	"os"

	"github.com/MuriLovin/fin-control/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file\n")
	}

	router := router.InitRouter()
	err = http.ListenAndServe(":"+os.Getenv("PORT"), router)

	if err != nil {
		log.Fatalf("Server closed with error: %s\n", err)
		return
	}
}
