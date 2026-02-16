package usecase

import (
	"context"
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/logger"
	"errors"
)

func (c *Core) StartBot(ctx context.Context, user domain.User) (msg *domain.Message, err error) {
	existed, err := c.fetchUserByID(ctx, user.ID)
	if err != nil {
		err = c.userRepo.Create(ctx, user)
		if err != nil {
			msgErr := "failed to create user"
			c.log.Log(logger.ERROR, msgErr, "id", user.ID.String(), "username", user.Username, "error", err.Error())
			return nil, errors.New(msgErr)
		}
		c.log.Log(logger.INFO, "user created", "id", user.ID.String(), "username", user.Username)
	} else {
		user = existed
		user.State = domain.UserStateStart
		err = c.userRepo.Update(ctx, user)
		if err != nil {
			msgErr := "failed to update user"
			c.log.Log(logger.ERROR, msgErr, "id", user.ID.String(), "username", user.Username)
			return nil, errors.New(msgErr)
		}
	}

	keyboard := domain.NewKeyboard()
	keyboard.Row().AddButton("Начать", domain.ActionOfferLeadMagnet)
	return domain.NewMessage(user.ChatID, `Привет 👋
	Ты попал в AI-систему контента и продаж.
	
	Здесь ты узнаешь:
	— почему AI-контент не приносит денег
	— как получать заявки без камеры
	— как собрать систему, а не хаос
	
	Готов начать?`, keyboard), nil
}
