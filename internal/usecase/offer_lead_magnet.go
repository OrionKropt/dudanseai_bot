package usecase

import (
	"context"
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/logger"
	"errors"
	"fmt"
	"net/url"
)

func (c *Core) OfferLeadMagnet(ctx context.Context, user domain.User) (msg *domain.Message, err error) {
	guideTitle := "Lead magnet 1"
	if user, err = c.fetchUserByID(ctx, user.ID); err != nil {
		return nil, err
	}
	user.State = domain.UserStateReady
	if err = c.userRepo.Update(ctx, user); err != nil {
		return nil, err

	}
	c.log.Log(logger.INFO, "offer lead magnet", "id", user.ID.String(), "username", user.Username)
	guide, err := c.materialRepo.FindMaterialByTitle(ctx, guideTitle)
	if err != nil {
		msgErr := fmt.Sprintf("Failed to get %s", guideTitle)
		c.log.Log(logger.ERROR, msgErr, "error", err.Error())
		return nil, errors.New(msgErr)
	}
	keyboard := domain.NewKeyboard()
	link, err := url.Parse(guide.URL.String())
	if err != nil {
		msgErr := "failed to parse URL"
		c.log.Log(logger.ERROR, msgErr, "error", err.Error())
		return nil, errors.New(msgErr)
	}
	keyboard.Row().AddButtonURL("📥 Забрать гайд ", link)
	return domain.NewMessage(user.ChatID, guide.Description, keyboard), nil
}
