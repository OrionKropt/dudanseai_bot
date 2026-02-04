package usecase

import (
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/logger"
	"errors"
	"strconv"
)

func (c *Core) StartBot(user domain.User) (msg *domain.Message, err error) {
	user.ID = domain.GenerateID(strconv.FormatInt(int64(user.ChatID), 10))
	_, ok := c.repo.GetUser(user.ID)
	if ok {
		c.log.Log(logger.INFO, "User already exists", "username", user.Username)
		return nil, errors.New("user already exists")
	}
	c.repo.CreateUser(user)
	c.log.Log(logger.INFO, "User created", "uuid", user.ID.String(), "username", user.Username)

	keyboard := domain.NewKeyboard()
	keyboard.Row().AddButton("Начать", domain.ActionQualification)
	return domain.NewMessage(user.ChatID, `Привет 👋
	Ты попал в AI-систему контента и продаж.
	
	Здесь ты узнаешь:
	— почему AI-контент не приносит денег
	— как получать заявки без камеры
	— как собрать систему, а не хаос
	
	Готов начать?`, keyboard), nil
}
