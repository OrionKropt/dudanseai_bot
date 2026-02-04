package main

import (
	"dudanseai_bot/configs"
	"dudanseai_bot/internal/app"
	"dudanseai_bot/pkg/logger"
	"log/slog"
	"os"
)

func initialize() (cfg *configs.Config, log *slog.Logger, err error) {
	cfg = configs.NewConfig()
	err = cfg.ReadConfig()
	if err != nil {
		return nil, nil, err
	}
	logHandler := logger.NewLogHandler(os.Stdout, logger.LogHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: logger.ParseLevel(cfg.LogLevel),
		},
	})
	log = slog.New(logHandler)
	return cfg, log, nil
}

func main() {
	cfg, log, err := initialize()
	if err != nil {
		panic(err)
	}
	log.Info("Initializing dudanseai_bot")
	app.Run(cfg, log)
}
