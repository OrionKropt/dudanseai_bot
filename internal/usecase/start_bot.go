package usecase

import (
	"context"
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/logger"
	"errors"
)

func (c *Core) startOnboarding(ctx context.Context, user domain.User) error {
	user.State = domain.UserStateOfferLeadMagnet
	err := c.userRepo.Create(ctx, user)
	if err != nil {
		msgErr := "failed to create user"
		c.log.Log(logger.ERROR, msgErr, "id", user.ID.String(), "username", user.Username, "error", err.Error())
		return errors.New(msgErr)
	}
	c.log.Log(logger.INFO, "user created", "id", user.ID.String(), "username", user.Username)
	return nil
}

func (c *Core) restartOnboarding(ctx context.Context, user domain.User) error {
	err := c.userRepo.Delete(ctx, user.ID)
	if err != nil {
		msgErr := "failed to delete user"
		c.log.Log(logger.ERROR, msgErr, "id", user.ID.String(), "username", user.Username, "error", err.Error())
		return errors.New(msgErr)
	}
	return c.startOnboarding(ctx, user)
}

func (c *Core) StartBot(ctx context.Context, user domain.User) (err error) {
	_, err = c.fetchUserByID(ctx, user.ID)
	if err != nil {
		if err := c.startOnboarding(ctx, user); err != nil {
			return err
		}
	} else {
		if err := c.restartOnboarding(ctx, user); err != nil {
			return err
		}
	}

	keyboard := domain.NewKeyboard()
	keyboard.Row().AddButton("Начать", domain.ActionOfferLeadMagnet)

	err = c.sender.SendMessage(domain.NewMessage(user.ChatID, `Привет 👋
Ты попал в AI-систему контента и продаж.

Здесь ты узнаешь:
	— почему AI-контент не приносит денег
	— как получать заявки без камеры
	— как собрать систему, а не хаос

Готов начать?`, keyboard))
	if err != nil {
		c.log.Log(logger.ERROR, "failed to start bot", "id", user.ID.String(), "username", user.Username, "error", err.Error())
	}
	return err
}
