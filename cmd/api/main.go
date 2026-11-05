package main

import (
	"inventory-api/internal/config"
	"inventory-api/internal/database"
	"log"
	"net/http"

)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.OpenConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database Connected Successfully")

	server := http.Server{
		Addr: ":" + cfg.AppPort,
	}

	log.Println("Server Running on port", cfg.AppPort)

	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	_ = db
}