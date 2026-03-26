package logger

import (
	"os"

	"go.uber.org/zap"
)

func New() (*zap.Logger, error) {
	if os.Getenv("APP_ENV") == "production" {
		return zap.NewProduction()
	}
	return zap.NewDevelopment()
}
