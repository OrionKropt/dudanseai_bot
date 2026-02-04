package app

import (
	"dudanseai_bot/configs"
	tgapi "dudanseai_bot/internal/controller/telegramapi"
	"dudanseai_bot/internal/infrastructure/repository"
	"dudanseai_bot/internal/usecase"
	"log/slog"
)

func Run(cfg *configs.Config, log *slog.Logger) {
	repo := repository.NewStorage()
	core := usecase.New(log, repo)
	bot := tgapi.NewBot(log, core, cfg.TelegramApiKey)
	err := bot.Init()
	if err != nil {
		log.Error("Bot initialization failed", "error", err.Error())
		return
	}
	bot.Start()
}
