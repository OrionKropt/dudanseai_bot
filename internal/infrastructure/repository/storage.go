package repository

import (
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/syncmap"
)

type Storage struct {
	Users syncmap.SyncMap[domain.UserID, domain.User]
}

func NewStorage() *Storage {
	return &Storage{Users: syncmap.NewSyncMap[domain.UserID, domain.User]()}
}

func (s *Storage) UpdateUser(u domain.User) {
	s.Users.Set(u.ID, u)
}

func (s *Storage) CreateUser(u domain.User) {
	s.Users.Set(u.ID, u)
}

func (s *Storage) DeleteUser(id domain.UserID) {
	if s.Users.Exists(id) {
		s.Users.Delete(id)
	}
}

func (s *Storage) GetUser(id domain.UserID) (domain.User, bool) {
	return s.Users.Get(id)
}
