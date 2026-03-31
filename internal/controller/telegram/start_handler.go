package telegram

import (
	"context"
	"dudanseai_bot/internal/domain"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (b *BotHandler) StartHandler(ctx context.Context, api *bot.Bot, update *models.Update) {
	chat := update.Message.Chat
	user := domain.NewUser(chat.ID, chat.Username, chat.FirstName, chat.LastName)
	b.uc.StartBot(ctx, user)
}
