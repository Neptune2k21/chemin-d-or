package app

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/neptune2k21/chemin-d-or/internal/config"
	"github.com/neptune2k21/chemin-d-or/internal/server"
	"github.com/neptune2k21/chemin-d-or/internal/storage"
	"github.com/neptune2k21/chemin-d-or/internal/tasks"
	"go.uber.org/zap"
)

type App struct {
	cfg    *config.Config
	db     *pgxpool.Pool
	logger *zap.Logger
	router http.Handler
}

func New(cfg *config.Config, logger *zap.Logger) (*App, error) {
	db, err := storage.NewPostgresPool(cfg.DatabaseURL())
	if err != nil {
		return nil, err
	}

	taskRepo := tasks.NewPostgresRepository(db)
	taskSvc := tasks.NewService(taskRepo)
	taskHandler := tasks.NewHandler(taskSvc)

	router := server.NewRouter(taskHandler)

	return &App{
		cfg:    cfg,
		db:     db,
		logger: logger,
		router: router,
	}, nil
}

func (a *App) Run() error {
	a.logger.Info("serveur démarré", zap.String("port", a.cfg.AppPort))
	return http.ListenAndServe(":"+a.cfg.AppPort, a.router)
}

func (a *App) Shutdown(_ context.Context) {
	a.db.Close()
}
