package telegram

import (
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/logger"
	"errors"
	"log/slog"

	"dudanseai_bot/internal/pkg/botapi"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type BotProvider struct {
	log logger.Logger
	api *botapi.BotAPI
}

func NewBotProvider(log *slog.Logger, api *botapi.BotAPI) BotProvider {
	return BotProvider{log: logger.Logger{Inst: log, Name: "Bot provider"}, api: api}
}

func (b *BotProvider) SendMessage(msg domain.Message) error {
	err := b.api.SendMessage(&bot.SendMessageParams{ChatID: msg.ChatID, Text: msg.Text, ReplyMarkup: b.api.CreateInlineMarkab(msg.Keys)})
	if err != nil {
		return errors.New("failed to send message")
	}
	return nil
}

func (b *BotProvider) SendVideoNote(note *domain.VideoNote) error {
	err := b.api.SendVideoNote(&bot.SendVideoNoteParams{ChatID: note.ChatID, VideoNote: &models.InputFileString{Data: note.File.TelegramFileID}})
	if err != nil {
		return errors.New("failed to send video note")
	}
	return nil
}
