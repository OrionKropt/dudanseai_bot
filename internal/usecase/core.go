package usecase

import (
	"context"
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/logger"
	"errors"
	"fmt"
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

	CreateUserMaterial(ctx context.Context, m domain.UserMaterial) error
	FindUserMaterialByMaterialID(ctx context.Context, userID domain.UserID, materialID domain.MaterialID) (um domain.UserMaterial, err error)
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
		c.log.Log(logger.INFO, "failed to fetch user by id", "id", id.String(), "error", err.Error())
		return domain.User{}, errors.New("user not exists")
	}
	return existed, nil
}

func (c *Core) fetchMaterialByTitle(ctx context.Context, title string) (domain.Material, error) {
	m, err := c.materialRepo.FindMaterialByTitle(ctx, title)
	if err != nil {
		msgErr := fmt.Sprintf("failed to get %s", title)
		c.log.Log(logger.ERROR, msgErr, "error", err.Error())
		return domain.Material{}, errors.New(msgErr)
	}
	return m, nil
}

func (c *Core) addMaterialToUser(ctx context.Context, u domain.User, m domain.Material) error {
	if _, err := c.materialRepo.FindUserMaterialByMaterialID(ctx, u.ID, m.ID); err == nil {
		return nil
	}

	if err := c.materialRepo.CreateUserMaterial(ctx, domain.CreateUserMaterial(u.ID, m.ID)); err != nil {
		msgErr := fmt.Sprintf("failed to add material %s to user %s", m.Title, u.Username)
		c.log.Log(logger.ERROR, msgErr, "error", err.Error())
		return errors.New(msgErr)
	}

	return nil
}
