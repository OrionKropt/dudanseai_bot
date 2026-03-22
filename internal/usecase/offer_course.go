package usecase

import (
	"context"
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/logger"
	"errors"
	"net/url"
)

func (c *Core) OfferCourse(ctx context.Context, user domain.User) (err error) {
	offerTitle := "cta_buy_course"
	if user, err = c.fetchUserByID(ctx, user.ID); err != nil {
		return err
	}
	c.log.Log(logger.INFO, "offer course", "id", user.ID.String(), "username", user.Username)
	courseAccessOffer, err := c.fetchMaterialByTitle(ctx, offerTitle)
	if err != nil {
		return err
	}

	keyboard := domain.NewKeyboard()
	link, err := url.Parse(courseAccessOffer.URL.String())
	if err != nil {
		msgErr := "failed to parse URL"
		c.log.Log(logger.ERROR, msgErr, "error", err.Error())
		return errors.New(msgErr)
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
		c.log.Log(logger.ERROR, "failed to offer course", "id", user.ID.String(), "username", user.Username, "error", err.Error())
	}
	return err
}
