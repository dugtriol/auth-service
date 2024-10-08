package app

import (
	`log/slog`
	`time`

	grpcapp `github.com/dugtriol/auth-service/internal/app/grpc`
	`github.com/dugtriol/auth-service/internal/services/auth`
	`github.com/dugtriol/auth-service/internal/storage/sqlite`
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, grpcPort, authService)

	return &App{
		GRPCServer: grpcApp,
	}
}
