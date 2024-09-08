package app

import (
	`log/slog`
	`time`

	grpcapp `github.com/dugtriol/auth-service/internal/app/grpc`
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	app := grpcapp.New(log, grpcPort)
	return &App{GRPCServer: app}
}
