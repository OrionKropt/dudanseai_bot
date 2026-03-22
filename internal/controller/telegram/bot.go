package telegram

import (
	"context"
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/internal/pkg/botapi"
	"dudanseai_bot/pkg/logger"
	"log/slog"

	"github.com/go-telegram/bot"
)

type UseCase interface {
	StartBot(context.Context, domain.User) error
	OfferLeadMagnet(context.Context, domain.User) error
	PresentSystem(ctx context.Context, user domain.User) error
	OfferCourse(ctx context.Context, user domain.User) error
}

type BotHandler struct {
	api *botapi.BotAPI
	log logger.Logger
	uc  UseCase
}

func NewBotHandler(inst *slog.Logger, api *botapi.BotAPI, uc UseCase) *BotHandler {
	return &BotHandler{log: logger.Logger{Inst: inst, Name: "Bot handler"}, api: api, uc: uc}
}

func (b *BotHandler) Init() {
	b.initHandlers()
}

func (b *BotHandler) initHandlers() {
	b.api.API.RegisterHandler(bot.HandlerTypeMessageText, "start", bot.MatchTypeCommand, b.StartHandler)
	b.api.RegisterActionHandler(domain.ActionOfferLeadMagnet, b.OfferLeadMagnetHandler)
}
