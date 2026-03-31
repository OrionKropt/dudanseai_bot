package app

import (
	"context"
	"dudanseai_bot/configs"
	tgctrl "dudanseai_bot/internal/controller/telegram"
	"dudanseai_bot/internal/infrastructure/repository/postgresql"
	tginfra "dudanseai_bot/internal/infrastructure/telegram"
	tgapi "dudanseai_bot/internal/pkg/botapi"
	"dudanseai_bot/internal/usecase"
	"log/slog"
)

func Run(cfg configs.Config, log *slog.Logger) {
	sqlClient, err := postgresql.NewClient(context.Background(), 5, cfg.DbBot)
	if err != nil {
		log.Error("Failed to create data base client", "error", err.Error())
		return
	}
	defer sqlClient.Close()

	userRepo, err := postgresql.NewUserRepository(&sqlClient)
	if err != nil {
		log.Error("Failed to create user repository", "error", err.Error())
		return
	}
	materialsRepo, err := postgresql.NewMaterialsRepository(&sqlClient)
	if err != nil {
		log.Error("Failed to create materials repository", "error", err.Error())
		return
	}

	log.Info("Repositories initialized")
	botClient := tgapi.NewBot(log, cfg.TelegramApiKey)
	err = botClient.Init(cfg.ProxyUrl)
	if err != nil {
		log.Error("Bot API initialization failed", "error", err.Error())
		return
	}

	botProvider := tginfra.NewBotProvider(log, botClient)
	core := usecase.New(log, &userRepo, &materialsRepo, &botProvider)

	botHandler := tgctrl.NewBotHandler(log, botClient, core)
	botHandler.Init()

	botClient.Start()
}
