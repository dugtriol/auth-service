package main

import (
	`fmt`
	`log/slog`
	`os`

	`github.com/dugtriol/auth-service/internal/app`
	`github.com/dugtriol/auth-service/internal/config`
)

const (
	path     = "./config/config.yaml"
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad(path)

	fmt.Println(cfg)

	log := setupLogger(envLocal)
	log.Info("starting application")
	// TODO: app initialization
	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)
	application.GRPCServer.MustRun()

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}

	return log
}
