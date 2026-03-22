package domain

import (
	"dudanseai_bot/internal/infrastructure/uuid"
	"time"
)

type UserMaterialID uuid.UUID

type UserMaterial struct {
	ID         UserMaterialID
	UserID     UserID
	MaterialID MaterialID
	SentAt     time.Time
}

func CreateUserMaterial(userID UserID, materialID MaterialID) UserMaterial {
	return UserMaterial{UserID: userID, MaterialID: materialID, SentAt: time.Now()}
}
