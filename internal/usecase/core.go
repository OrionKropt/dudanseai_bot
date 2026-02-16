package usecase

import (
	"context"
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/logger"
	"errors"
	"log/slog"
)

type UserRepository interface {
	Create(ctx context.Context, u domain.User) error
	Delete(ctx context.Context, id domain.UserID) error
	FindOne(ctx context.Context, id domain.UserID) (domain.User, error)
	Update(ctx context.Context, u domain.User) error
}

type MaterialsRepository interface {
	CreateMaterial(ctx context.Context, m domain.Material) error
	DeleteMaterial(ctx context.Context, id domain.MaterialID) error
	FindOneMaterial(ctx context.Context, id domain.MaterialID) (domain.Material, error)
	UpdateMaterial(ctx context.Context, m domain.Material) error
	FindMaterialByTitle(ctx context.Context, title string) (domain.Material, error)
	// TODO LoadFile ...
}

type Core struct {
	userRepo     UserRepository
	materialRepo MaterialsRepository
	log          logger.Logger
}

func New(log *slog.Logger, ur UserRepository, mr MaterialsRepository) *Core {
	return &Core{userRepo: ur, materialRepo: mr, log: logger.Logger{Name: "Core", Inst: log}}
}

func (c *Core) fetchUserByID(ctx context.Context, id domain.UserID) (domain.User, error) {
	existed, err := c.userRepo.FindOne(ctx, id)
	if err != nil {
		c.log.Log(logger.INFO, "failed to fetch user by id", "id", id, "err", err.Error())
		return domain.User{}, errors.New("user not exists")
	}
	return existed, nil
}
