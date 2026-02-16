package telegramapi

import (
	"context"
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/logger"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (b *BotAPI) ButtonCallbackHandler(ctx context.Context, api *bot.Bot, msg models.MaybeInaccessibleMessage, data []byte) {
	user := domain.NewUser(msg.Message.Chat.ID, msg.Message.Chat.Username, msg.Message.Chat.FirstName, msg.Message.Chat.LastName)
	actionType := domain.ActionType(data)
	err := b.actionsHandler.Action(ctx, user, actionType)
	if err != nil {
		b.log.Log(logger.ERROR, "Failed to call action", "action type", actionType, "error", err.Error())
	}
}
