package main

import (
	"context"
	"log"

	"github.com/neptune2k21/chemin-d-or/internal/app"
	"github.com/neptune2k21/chemin-d-or/internal/config"
	"github.com/neptune2k21/chemin-d-or/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("erreur lors du chargement de la configuration: %v", err)
	}

	l, err := logger.New()
	if err != nil {
		log.Fatalf("erreur lors de l'initialisation du logger: %v", err)
	}
	defer l.Sync() //nolint:errcheck

	a, err := app.New(cfg, l)
	if err != nil {
		l.Fatal("erreur lors de l'initialisation de l'application", zap.Error(err))
	}
	defer a.Shutdown(context.Background())

	if err := a.Run(); err != nil {
		l.Fatal("serveur arrêté", zap.Error(err))
	}
}
