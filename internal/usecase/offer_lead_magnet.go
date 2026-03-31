package usecase

import (
	"context"
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/logger"
	"net/url"
	"time"
)

func (c *Core) OfferLeadMagnet(ctx context.Context, user domain.User) {
	var err error
	defer func() {
		if err != nil {
			c.logError(user, "failed to offer lead magnet", err)
		}
	}()
	guideTitle := "lead_magnet_1"
	c.log.Log(logger.INFO, "offer lead magnet", "id", user.ID.String(), "username", user.Username)
	if user, err = c.fetchUserByID(ctx, user.ID); err != nil {
		return
	}
	guide, err := c.fetchMaterialByTitle(ctx, guideTitle)
	if err != nil {
		return
	}

	keyboard := domain.NewKeyboard()
	link, err := url.Parse(guide.URL.String())
	if err != nil {
		return
	}
	keyboard.Row().AddButtonURL("📥 Забрать гайд", link)

	if err = c.addMaterialToUser(ctx, user, guide); err != nil {
		return
	}

	user.State = domain.UserStatePresentSystem
	user.LastInteraction = time.Now()
	if err = c.userRepo.Update(ctx, user); err != nil {
		return
	}

	if err = c.sender.SendMessage(domain.NewMessage(user.ChatID, guide.Description, keyboard)); err != nil {
		return
	}

	detachedCtx := context.WithoutCancel(ctx)
	time.AfterFunc(time.Second*24, func() {
		err = c.PresentSystem(detachedCtx, user)
		c.logError(user, "failed to present system", err)
		time.AfterFunc(time.Second*24, func() {
			err = c.OfferCourse(detachedCtx, user)
			c.logError(user, "failed to offer course", err)
		})
	})
}
