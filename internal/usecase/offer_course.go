package usecase

import (
	"context"
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/logger"
	"fmt"
	"net/url"
)

func (c *Core) OfferCourse(ctx context.Context, user domain.User) (err error) {
	offerTitle := "cta_buy_course"
	c.log.Log(logger.INFO, "offer course", "id", user.ID.String(), "username", user.Username)
	if user, err = c.fetchUserByID(ctx, user.ID); err != nil {
		return err
	}

	courseAccessOffer, err := c.fetchMaterialByTitle(ctx, offerTitle)
	if err != nil {
		return err
	}

	keyboard := domain.NewKeyboard()
	link, err := url.Parse(courseAccessOffer.URL.String())
	if err != nil {
		return err
	}
	keyboard.Row().AddButtonURL("🚀 Получить полный курс", link)

	if err := c.addMaterialToUser(ctx, user, courseAccessOffer); err != nil {
		return err
	}

	user.State = domain.UserStateReady
	if err := c.userRepo.Update(ctx, user); err != nil {
		return err
	}

	err = c.sender.SendMessage(domain.NewMessage(user.ChatID, courseAccessOffer.Description, keyboard))
	if err != nil {
		err = fmt.Errorf("failed to send message: %v", err)
	}
	return err
}
