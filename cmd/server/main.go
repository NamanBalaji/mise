package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/NamanBalaji/mise/internal/config"
	"github.com/NamanBalaji/mise/internal/handlers"
)

func main() {

	log.Println("Starting Mise")
	app := &config.AppConfig{
		Version: 1,
	}

	repo := handlers.NewRepo(app)
	handlers.NewHandlers(repo)

	port := flag.String("p", "6379", "")
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
