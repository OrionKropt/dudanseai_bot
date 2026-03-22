package usecase

import (
	"context"
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/logger"
	"errors"
	"net/url"
	"time"
)

func (c *Core) OfferLeadMagnet(ctx context.Context, user domain.User) (err error) {
	guideTitle := "lead_magnet_1"
	if user, err = c.fetchUserByID(ctx, user.ID); err != nil {
		return err
	}
	c.log.Log(logger.INFO, "offer lead magnet", "id", user.ID.String(), "username", user.Username)

	guide, err := c.fetchMaterialByTitle(ctx, guideTitle)
	if err != nil {
		return err
	}
	keyboard := domain.NewKeyboard()
	link, err := url.Parse(guide.URL.String())
	if err != nil {
		msgErr := "failed to parse URL"
		c.log.Log(logger.ERROR, msgErr, "error", err.Error())
		return errors.New(msgErr)
	}
	keyboard.Row().AddButtonURL("📥 Забрать гайд", link)

	if err := c.addMaterialToUser(ctx, user, guide); err != nil {
		return err
	}

	user.State = domain.UserStatePresentSystem
	user.LastInteraction = time.Now()
	if err := c.userRepo.Update(ctx, user); err != nil {
		return err
	}

	if err := c.sender.SendMessage(domain.NewMessage(user.ChatID, guide.Description, keyboard)); err != nil {
		return err
	}

	detachedCtx := context.WithoutCancel(ctx)
	time.AfterFunc(time.Second*24, func() {
		_ = c.PresentSystem(detachedCtx, user) // TODO don't ignore error

		time.AfterFunc(time.Second*24, func() {
			_ = c.OfferCourse(detachedCtx, user)
		})
	})

	return nil
}
