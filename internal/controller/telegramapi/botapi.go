package telegramapi

import (
	"context"
	"dudanseai_bot/internal/controller"
	"dudanseai_bot/internal/controller/telegramapi/middleware"
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/pkg/logger"
	"errors"
	"log/slog"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/keyboard/inline"
)

type UseCase interface {
	StartBot(context.Context, domain.User) (*domain.Message, error)
	OfferLeadMagnet(context.Context, domain.User) (*domain.Message, error)
	PresentSystem(ctx context.Context, user domain.User) (*domain.VideoNote, error)
	OfferCourse(ctx context.Context, user domain.User) (*domain.Message, error)
}

type BotAPI struct {
	API            *bot.Bot
	log            *logger.Logger
	uc             UseCase
	actionsHandler *controller.ActionHandler
	ctx            context.Context
	cancel         context.CancelFunc
	token          string
}

func NewBot(inst *slog.Logger, uc UseCase, token string) *BotAPI {
	return &BotAPI{
		token:          token,
		log:            &logger.Logger{Inst: inst, Name: "BotAPI"},
		actionsHandler: controller.NewActionHandler(),
		uc:             uc,
	}
}

func (b *BotAPI) initHandlers() {
	b.API.RegisterHandler(bot.HandlerTypeMessageText, "start", bot.MatchTypeCommand, b.StartHandler)
	b.actionsHandler.RegisterHandler(domain.ActionOfferLeadMagnet, b.OfferLeadMagnetHandler)
}

func (b *BotAPI) Init() (err error) {
	b.log.Log(logger.INFO, "Initializing BotAPI")

	opts := []bot.Option{
		bot.WithMiddlewares(middleware.LoggingMiddleware(b.log)),
		bot.WithDefaultHandler(func(ctx context.Context, bot *bot.Bot, update *models.Update) {
			b.log.Log(logger.INFO, "Default handler")
		}),
	}

	b.API, err = bot.New(b.token, opts...)
	if err != nil {
		return err
	}
	b.initHandlers()
	return nil
}

func (b *BotAPI) Start() {
	b.log.Log(logger.INFO, "Bot start")
	b.ctx, b.cancel = signal.NotifyContext(context.Background(), os.Interrupt)
	// b.ctx, b.cancel = context.WithCancel(context.Background())
	b.API.Start(b.ctx)
}

func (b *BotAPI) Stop() {
	b.log.Log(logger.INFO, "Bot stop")
	b.cancel()
}

func (b *BotAPI) sendMessage(params *bot.SendMessageParams) error {
	_, err := b.API.SendMessage(b.ctx, params)
	return err
}

func (b *BotAPI) LoggingErrors(err error) {
	var tmrErr *bot.TooManyRequestsError
	if errors.As(err, &tmrErr) {
		b.log.Log(logger.ERROR, "Error 429", "error", err.Error())
		b.log.Log(logger.ERROR, "Received TooManyRequestsError with retry_after:", tmrErr.RetryAfter)
		return
	}

	switch {
	case errors.Is(err, bot.ErrorForbidden):
		b.log.Log(logger.ERROR, "Error 403", "error", err.Error())
	case errors.Is(err, bot.ErrorBadRequest):
		b.log.Log(logger.ERROR, "Error 400", "error", err.Error())
	case errors.Is(err, bot.ErrorUnauthorized):
		b.log.Log(logger.ERROR, "Error 401", "error", err.Error())
	case errors.Is(err, bot.ErrorNotFound):
		b.log.Log(logger.ERROR, "Error 404", "error", err.Error())
	case errors.Is(err, bot.ErrorConflict):
		b.log.Log(logger.ERROR, "Error 409", "msg", err.Error())
	default:
		b.log.Log(logger.ERROR, "Unknown error", "msg", err.Error())
	}
}

func (b *BotAPI) CreateInlineMarkab(keyboard domain.Keyboard) (kb *inline.Keyboard) {
	markup := keyboard.Markup()
	kb = inline.New(b.API)
	for _, r := range markup {
		kb = kb.Row()
		for _, btn := range r {
			if btn.URL != nil {
				kb.ButtonURL(btn.Text, *btn.URL)
			} else {
				kb = kb.Button(btn.Text, []byte(btn.Action), b.ButtonCallbackHandler)
			}
		}
	}
	return kb
}
