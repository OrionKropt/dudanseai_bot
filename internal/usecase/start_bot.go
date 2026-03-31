package usecase

import (
	"context"
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/logger"
	"fmt"
)

func (c *Core) startOnboarding(ctx context.Context, user domain.User) error {
	user.State = domain.UserStateOfferLeadMagnet
	err := c.userRepo.Create(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to crate user: %v", err)
	}
	c.log.Log(logger.INFO, "user created", "id", user.ID.String(), "username", user.Username)
	return nil
}

func (c *Core) restartOnboarding(ctx context.Context, user domain.User) error {
	err := c.userRepo.Delete(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return c.startOnboarding(ctx, user)
}

func (c *Core) StartBot(ctx context.Context, user domain.User) {
	var err error
	defer func() {
		if err != nil {
			c.logError(user, "failed to start bot", err)
		}
	}()

	c.log.Log(logger.INFO, "start bot", "id", user.ID.String(), "username", user.Username)
	_, err = c.fetchUserByID(ctx, user.ID)
	if err != nil {
		if e := c.startOnboarding(ctx, user); e != nil {
			err = fmt.Errorf("failed to start onboarding: %v", e)
			return
		}
	} else {
		if e := c.restartOnboarding(ctx, user); e != nil {
			err = fmt.Errorf("failed to restart onboarding: %v", e)
			return
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
		err = fmt.Errorf("failed to send message: %v", err)
	}
}
