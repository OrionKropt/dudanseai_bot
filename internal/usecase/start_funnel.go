package usecase

import (
	"context"
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/logger"
	"errors"
)

func (c *Core) StartFunnel(ctx context.Context, user domain.User) (msg *domain.Message, err error) {
	if user, err = c.fetchUserByID(ctx, user.ID); err != nil {
		return nil, err
	}

	user.State = domain.UserStateQualification
	err = c.userRepo.Update(ctx, user)
	if err != nil {
		msgErr := "failed to update user state"
		c.log.Log(logger.ERROR, msgErr, "id", user.ID.String(), "error", err.Error())
		return nil, errors.New(msgErr)
	}
	c.log.Log(logger.INFO, "User started funnel", "uuid", user.ID.String(), "username", user.Username)

	keyboard := domain.NewKeyboard()
	keyboard.Row().AddButton("🟣 SMM / маркетолог", domain.ActionOfferLeadMagnet)
	keyboard.Row().AddButton("🔵 Бизнес / предприниматель", domain.ActionOfferLeadMagnet)
	keyboard.Row().AddButton("🟢 Стартап / проект", domain.ActionOfferLeadMagnet)
	return domain.NewMessage(user.ChatID, "Чтобы дать тебе максимум пользы — подскажи, кто ты?", keyboard), nil
}
