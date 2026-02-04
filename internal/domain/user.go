package domain

import (
	"dudanseai_bot/internal/infrastructure/uuid"
	"time"
)

type ActionType string

const (
	ActionStartFunnel     = "START_FUNNEL"
	ActionGetGuide        = "GET_GUIDE"
	ActionOfferLeadMagnet = "OFFER_LEAD_MAGNET"
	ActionQualification   = "QUALIFICATION"
)

const (
	UserStateStart             = "START"
	UserStateQualification     = "QUALIFICATION"
	UserStateGettingLeadMagnet = "GETTING_LEAD_MAGNET"
	UserStateReady             = "READY"
)

const (
	UserRoleDefault  = "user"
	UserRoleAdmin    = "admin"
	UserRoleEngineer = "engineer"
)

func GenerateID(username string) UserID {
	id := uuid.Generate([]byte(username))
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
}

func NewUser(chatID int64, username, name, surname string) User {
	return User{
		ChatID:          ChatID(chatID),
		Username:        username,
		Name:            name,
		Surname:         surname,
		Role:            UserRoleDefault,
		State:           UserStateStart,
		LastInteraction: time.Now(),
	}
}
