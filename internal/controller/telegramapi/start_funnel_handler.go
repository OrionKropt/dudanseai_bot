package telegramapi

import (
	"dudanseai_bot/internal/domain"

	"github.com/go-telegram/bot"
)

func (b *BotAPI) StartFunnelHandler(user domain.User) {
	msg, err := b.uc.StartFunnel(user)
	if err != nil {
		return
	}

	err = b.sendMessage(&bot.SendMessageParams{ChatID: msg.ChatID, Text: msg.Text, ReplyMarkup: b.CreateInlineMarkab(msg.Keys)})
	if err != nil {
		b.LoggingErrors(err)
	}
}
