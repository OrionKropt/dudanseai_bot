package telegramapi

import (
	"context"
	"dudanseai_bot/internal/domain"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (b *BotAPI) OfferLeadMagnetHandler(ctx context.Context, user domain.User) {
	msg, err := b.uc.OfferLeadMagnet(ctx, user)
	if err != nil {
		return
	}

	err = b.sendMessage(&bot.SendMessageParams{ChatID: user.ChatID, Text: msg.Text, ReplyMarkup: b.CreateInlineMarkab(msg.Keys)})
	if err != nil {
		b.LoggingErrors(err)
	}

	detachedCtx := context.WithoutCancel(ctx)
	time.AfterFunc(time.Hour*24, func() {
		note, err := b.uc.PresentSystem(detachedCtx, user)
		if err == nil {
			err = b.sendVideoNote(&bot.SendVideoNoteParams{ChatID: user.ChatID, VideoNote: &models.InputFileString{Data: note.File.TelegramFileID}})
			if err != nil {
				b.LoggingErrors(err)
			}
		}
		
		time.AfterFunc(time.Hour*24, func() {
			msg, err = b.uc.OfferCourse(detachedCtx, user)
			if err != nil {
				return
			}
			err = b.sendMessage(&bot.SendMessageParams{ChatID: user.ChatID, Text: msg.Text, ReplyMarkup: b.CreateInlineMarkab(msg.Keys)})
			if err != nil {
				b.LoggingErrors(err)
			}
		})
	})
}
