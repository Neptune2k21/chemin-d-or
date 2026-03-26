package main

import (
	"log"
	"net/http"

	"github.com/neptune2k21/chemin-d-or/internal/config"
	"github.com/neptune2k21/chemin-d-or/internal/server"
	"github.com/neptune2k21/chemin-d-or/internal/storage"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := storage.NewPostgresPool(cfg.DatabaseURL())
	if err != nil {
		log.Fatalf("erreur lors de la connexion à la base de données: %v", err)
	}
	defer db.Close()

	router := server.NewRouter()

	log.Printf("serveur demarrer sur  :%s", cfg.AppPort)
	if err := http.ListenAndServe(":"+cfg.AppPort, router); err != nil {
		log.Fatalf("erreur lors du démarrage du serveur: %v", err)
	}
}
