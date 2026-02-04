package usecase

import (
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/logger"
	"errors"
	"log/slog"
	"strconv"
)

type Repository interface {
	CreateUser(u domain.User)
	DeleteUser(id domain.UserID)
	GetUser(id domain.UserID) (domain.User, bool)
	UpdateUser(u domain.User)
}
type Core struct {
	repo Repository
	log  *logger.Logger
}

func New(log *slog.Logger, repo Repository) *Core {
	return &Core{repo: repo, log: &logger.Logger{Name: "Core", Inst: log}}
}

func (c *Core) fetchUserByChatID(chatID domain.ChatID) (domain.User, error) {
	id := domain.GenerateID(strconv.FormatInt(int64(chatID), 10))
	existed, ok := c.repo.GetUser(id)
	if !ok {
		c.log.Log(logger.INFO, "User not exists")
		return domain.User{}, errors.New("user not exists")
	}
	return existed, nil
}
