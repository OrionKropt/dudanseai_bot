package telegram

import (
	"context"
	"dudanseai_bot/internal/domain"
)

func (b *BotHandler) OfferLeadMagnetHandler(ctx context.Context, user domain.User) {
	b.uc.OfferLeadMagnet(ctx, user)
}
