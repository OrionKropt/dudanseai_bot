package repository

import (
	"context"
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/syncmap"
	"errors"
)

type Storage struct {
	Users syncmap.SyncMap[domain.UserID, domain.User]
}

func NewStorage() Storage {
	return Storage{Users: syncmap.NewSyncMap[domain.UserID, domain.User]()}
}

func (s *Storage) UpdateUser(ctx context.Context, u domain.User) error {
	s.Users.Set(u.ID, u)
	return nil
}

func (s *Storage) CreateUser(ctx context.Context, u domain.User) error {
	s.Users.Set(u.ID, u)
	return nil
}

func (s *Storage) DeleteUser(ctx context.Context, id domain.UserID) error {
	if s.Users.Exists(id) {
		s.Users.Delete(id)
	} else {
		return errors.New("user not found")
	}
	return nil
}

func (s *Storage) GetUser(ctx context.Context, id domain.UserID) (domain.User, error) {
	user, ok := s.Users.Get(id)
	if !ok {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}
