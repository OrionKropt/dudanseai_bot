package app

import (
	"context"
	"dudanseai_bot/configs"
	tgapi "dudanseai_bot/internal/controller/telegramapi"
	"dudanseai_bot/internal/infrastructure/repository/postgresql"
	"dudanseai_bot/internal/usecase"
	"log/slog"
)

func Run(cfg configs.Config, log *slog.Logger) {
	client, err := postgresql.NewClient(context.Background(), 5, cfg.DbBot, log)
	if err != nil {
		log.Error("Failed to create data base client", "error", err.Error())
		return
	}
	defer client.Close()

	userRepo, err := postgresql.NewUserRepository(&client)
	if err != nil {
		log.Error("Failed to create user repository", "error", err.Error())
		return
	}
	materialsRepo, err := postgresql.NewMaterialsRepository(&client)
	if err != nil {
		log.Error("Failed to create materials repository", "error", err.Error())
		return
	}

	log.Info("Repositories initialized")
	core := usecase.New(log, &userRepo, &materialsRepo)
	bot := tgapi.NewBot(log, core, cfg.TelegramApiKey)

	err = bot.Init(cfg.ProxyUrl)
	if err != nil {
		log.Error("Bot initialization failed", "error", err.Error())
		return
	}
	bot.Start()
}
