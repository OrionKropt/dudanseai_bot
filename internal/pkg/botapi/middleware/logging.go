package middleware

import (
	"context"
	"dudanseai_bot/pkg/logger"
	"fmt"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func LoggingMiddleware(log *logger.Logger) func(next bot.HandlerFunc) bot.HandlerFunc {
	return func(next bot.HandlerFunc) bot.HandlerFunc {
		return func(ctx context.Context, b *bot.Bot, update *models.Update) {
			msg := update.Message
			if msg == nil {
				log.Log(logger.INFO, "empty message")
			} else {
				log.Log(logger.INFO, fmt.Sprintf("message %s from %s", msg.Text, msg.From.Username))
				if update.Message.VideoNote != nil {
					fileID := update.Message.VideoNote.FileID
					log.Log(logger.INFO, fmt.Sprintf("video file id: %s", fileID))
				}
			}

			next(ctx, b, update)
		}
	}
}
