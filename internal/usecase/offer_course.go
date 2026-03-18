package usecase

import (
	"context"
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/logger"
	"errors"
	"fmt"
	"net/url"
)

func (c *Core) OfferCourse(ctx context.Context, user domain.User) (msg *domain.Message, err error) {
	offerTitle := "cta_buy_course"
	if user, err = c.fetchUserByID(ctx, user.ID); err != nil {
		return nil, err
	}
	c.log.Log(logger.INFO, "offer course", "id", user.ID.String(), "username", user.Username)
	courseAccessOffer, err := c.materialRepo.FindMaterialByTitle(ctx, offerTitle)
	if err != nil {
		msgErr := fmt.Sprintf("failed to get %s", offerTitle)
		c.log.Log(logger.ERROR, msgErr, "error", err.Error())
		return nil, errors.New(msgErr)
	}
	keyboard := domain.NewKeyboard()
	link, err := url.Parse(courseAccessOffer.URL.String())
	if err != nil {
		msgErr := "failed to parse URL"
		c.log.Log(logger.ERROR, msgErr, "error", err.Error())
		return nil, errors.New(msgErr)
	}
	keyboard.Row().AddButtonURL("🚀 Получить полный курс", link)

	user.State = domain.UserStateReady
	if err := c.userRepo.Update(ctx, user); err != nil {
		return nil, err

	}

	return domain.NewMessage(user.ChatID, courseAccessOffer.Description, keyboard), nil
}
