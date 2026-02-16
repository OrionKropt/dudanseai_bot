package domain

import (
	"dudanseai_bot/internal/infrastructure/uuid"
	"strconv"
	"time"
)

type ActionType string

const (
	ActionOfferLeadMagnet = "OFFER_LEAD_MAGNET"
)

const (
	UserStateStart = "START"
	UserStateReady = "READY"
)

const (
	UserRoleDefault  = "user"
	UserRoleAdmin    = "admin"
	UserRoleEngineer = "engineer"
)

func generateUserID(chatID int64) UserID {
	id := uuid.Generate([]byte(strconv.FormatInt(chatID, 10)))
	return UserID(id)
}

type UserID uuid.UUID

func (u UserID) String() string {
	return uuid.UUID(u).String()
}

type ChatID int64
type UserState string
type User struct {
	ID              UserID
	ChatID          ChatID
	Name            string
	Surname         string
	Username        string
	Role            string
	State           UserState
	LastInteraction time.Time
	CreatedAt       time.Time
}

func NewUser(chatID int64, username, name, surname string) User {
	return User{
		ID:              generateUserID(chatID),
		ChatID:          ChatID(chatID),
		Username:        username,
		Name:            name,
		Surname:         surname,
		Role:            UserRoleDefault,
		State:           UserStateStart,
		LastInteraction: time.Now(),
	}
}
