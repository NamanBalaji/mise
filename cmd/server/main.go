package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/NamanBalaji/mise/internal/config"
	"github.com/NamanBalaji/mise/internal/database"
	"github.com/NamanBalaji/mise/internal/handlers"
)

func main() {

	log.Println("Starting Mise")
	app := &config.AppConfig{
		Version: 1,
	}

	port := flag.String("p", "6379", "The port you want to run the server on")
	memory := flag.Bool("m", true, "Set to false if you want data to persist")
	flag.Parse()

	// Initialize the database
	db := database.NewDB(*memory)

	// Initialize the handlers
	repo := handlers.NewRepo(app, db)
	handlers.NewHandlers(repo)

	portNum := fmt.Sprintf(":%s", *port)

	srv := &http.Server{
		Addr:    portNum,
		Handler: routes(),
	}

	log.Println("Server running on port", portNum)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
