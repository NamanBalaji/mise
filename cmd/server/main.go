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

	db := database.NewDB(true)
	repo := handlers.NewRepo(app, db)
	handlers.NewHandlers(repo)

	port := flag.String("p", "6379", "")
	flag.Parse()

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
