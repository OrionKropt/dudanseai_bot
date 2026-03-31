package usecase

import (
	"context"
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/logger"
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

type MessageSender interface {
	SendMessage(msg domain.Message) error
	SendVideoNote(note *domain.VideoNote) error
}

type Core struct {
	userRepo     UserRepository
	materialRepo MaterialsRepository
	sender       MessageSender
	log          logger.Logger
}

func New(log *slog.Logger, ur UserRepository, mr MaterialsRepository, s MessageSender) *Core {
	return &Core{userRepo: ur, materialRepo: mr, sender: s, log: logger.Logger{Name: "Core", Inst: log}}
}

func (c *Core) fetchUserByID(ctx context.Context, id domain.UserID) (domain.User, error) {
	existed, err := c.userRepo.FindOne(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to fetch user by id: %s: %v", id.String(), err)
	}
	return existed, nil
}

func (c *Core) fetchMaterialByTitle(ctx context.Context, title string) (domain.Material, error) {
	m, err := c.materialRepo.FindMaterialByTitle(ctx, title)
	if err != nil {
		return domain.Material{}, fmt.Errorf("failed to find material by title: %s: %v", title, err)
	}
	return m, nil
}

func (c *Core) addMaterialToUser(ctx context.Context, u domain.User, m domain.Material) error {
	if _, err := c.materialRepo.FindUserMaterialByMaterialID(ctx, u.ID, m.ID); err == nil {
		return nil
	}

	if err := c.materialRepo.CreateUserMaterial(ctx, domain.CreateUserMaterial(u.ID, m.ID)); err != nil {
		return fmt.Errorf("failed to add material %s to user %s: %v", m.Title, u.Username, err)
	}

	return nil
}

func (c *Core) logError(u domain.User, msg string, err error) {
	c.log.Log(logger.ERROR, "failed to offer lead magnet", "id", u.ID.String(),
		"username", u.Username, "msg", msg, "error", err.Error())
}
