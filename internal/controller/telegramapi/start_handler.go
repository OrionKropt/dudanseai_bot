package telegramapi

import (
	"context"
	"dudanseai_bot/internal/domain"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (b *BotAPI) StartHandler(ctx context.Context, api *bot.Bot, update *models.Update) {
	var err error
	chat := update.Message.Chat
	user := domain.NewUser(chat.ID, chat.Username, chat.FirstName, chat.LastName)
	msg, err := b.uc.StartBot(user)
	if err != nil {
		return
	}

	err = b.sendMessage(&bot.SendMessageParams{ChatID: msg.ChatID, Text: msg.Text, ReplyMarkup: b.CreateInlineMarkab(msg.Keys)})
	if err != nil {
		b.LoggingErrors(err)
	}
}
