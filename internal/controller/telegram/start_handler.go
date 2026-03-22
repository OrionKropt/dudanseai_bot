package telegram

import (
	"context"
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/logger"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (b *BotHandler) StartHandler(ctx context.Context, api *bot.Bot, update *models.Update) {
	chat := update.Message.Chat
	user := domain.NewUser(chat.ID, chat.Username, chat.FirstName, chat.LastName)
	if err := b.uc.StartBot(ctx, user); err != nil {
		b.log.Log(logger.ERROR, "failed to start api", "id", user.ID.String(), "username", chat.Username, "error", err.Error())
	}
}
