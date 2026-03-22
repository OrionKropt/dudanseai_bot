package usecase

import (
	"context"
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/logger"
)

func (c *Core) PresentSystem(ctx context.Context, user domain.User) (note *domain.VideoNote, err error) {
	c.log.Log(logger.INFO, "present system", "id", user.ID.String(), "username", user.Username)
	videoNoteTitle := "offer_system_circle"
	if user, err = c.fetchUserByID(ctx, user.ID); err != nil {
		return nil, err
	}

	videoNote, err := c.fetchMaterialByTitle(ctx, videoNoteTitle)
	if err != nil {
		return nil, err
	}

	if err := c.addMaterialToUser(ctx, user, videoNote); err != nil {
		c.log.Log(logger.ERROR, err.Error())
		return nil, err
	}

	user.State = domain.UserStateOfferCourse
	if err := c.userRepo.Update(ctx, user); err != nil {
		return nil, err

	}

	return domain.NewVideoNote(user.ChatID, domain.NewFileByID(videoNote.TelegramFileID)), nil
}
