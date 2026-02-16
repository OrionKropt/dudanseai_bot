package telegramapi

import (
	"context"
	"dudanseai_bot/internal/domain"

	"github.com/go-telegram/bot"
)

func (b *BotAPI) OfferLeadMagnetHandler(ctx context.Context, user domain.User) {
	msg, err := b.uc.OfferLeadMagnet(ctx, user)
	if err != nil {
		return
	}

	err = b.sendMessage(&bot.SendMessageParams{ChatID: msg.ChatID, Text: msg.Text, ReplyMarkup: b.CreateInlineMarkab(msg.Keys)})
	if err != nil {
		b.LoggingErrors(err)
	}
}
