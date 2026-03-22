package telegram

import (
	"context"
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/logger"
)

func (b *BotHandler) OfferLeadMagnetHandler(ctx context.Context, user domain.User) {
	if err := b.uc.OfferLeadMagnet(ctx, user); err != nil {
		b.log.Log(logger.ERROR, "failed to offer lead magnet", "id", user.ID.String(), "username", user.Username, "error", err.Error())
	}
}
