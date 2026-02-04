package usecase

import (
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/logger"
)

func (c *Core) StartFunnel(user domain.User) (msg *domain.Message, err error) {
	if user, err = c.fetchUserByChatID(user.ChatID); err != nil {
		return nil, err
	}

	user.State = domain.UserStateQualification
	c.repo.UpdateUser(user)
	c.log.Log(logger.INFO, "User started funnel", "uuid", user.ID.String(), "username", user.Username)

	keyboard := domain.NewKeyboard()
	keyboard.Row().AddButton("🟣 SMM / маркетолог", domain.ActionOfferLeadMagnet)
	keyboard.Row().AddButton("🔵 Бизнес / предприниматель", domain.ActionOfferLeadMagnet)
	keyboard.Row().AddButton("🟢 Стартап / проект", domain.ActionOfferLeadMagnet)
	return domain.NewMessage(user.ChatID, "Чтобы дать тебе максимум пользы — подскажи, кто ты?", keyboard), nil
}
