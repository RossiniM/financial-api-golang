package main

import (
	"context"
	"net/http"

	financial "example.com/m/v2/internal/domain/financial/module"
	logger "example.com/m/v2/internal/infra/logger"
	server "example.com/m/v2/internal/infra/server"

	"github.com/go-chi/chi/v5"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		server.Module,
		logger.Module,
		financial.Module,
		fx.Invoke(registerHooks),
	).Run()
}

func registerHooks(
	lifecycle fx.Lifecycle, router chi.Router, logger *zap.SugaredLogger,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				logger.Info("Listening on localhost:8080")
				go http.ListenAndServe(":8080", router)
				return nil
			},
			OnStop: func(context.Context) error {
				return logger.Sync()
			},
		},
	)
}
