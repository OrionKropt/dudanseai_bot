package usecase

import (
	"context"
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/logger"
	"fmt"
)

func (c *Core) PresentSystem(ctx context.Context, user domain.User) (err error) {
	c.log.Log(logger.INFO, "present system", "id", user.ID.String(), "username", user.Username)
	videoNoteTitle := "offer_system_circle"
	if user, err = c.fetchUserByID(ctx, user.ID); err != nil {
		return err
	}

	videoNote, err := c.fetchMaterialByTitle(ctx, videoNoteTitle)
	if err != nil {
		return err
	}

	if err := c.addMaterialToUser(ctx, user, videoNote); err != nil {
		return err
	}

	user.State = domain.UserStateOfferCourse
	if err := c.userRepo.Update(ctx, user); err != nil {
		return err

	}

	err = c.sender.SendVideoNote(domain.NewVideoNote(user.ChatID, domain.NewFileByID(videoNote.TelegramFileID)))
	if err != nil {
		err = fmt.Errorf("failed to send video note: %v", err)
	}
	return err
}
