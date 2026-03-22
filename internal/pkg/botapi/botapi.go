package botapi

import (
	"context"
	"dudanseai_bot/internal/domain"
	"dudanseai_bot/internal/pkg/botapi/middleware"
	"dudanseai_bot/pkg/logger"
	"errors"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/keyboard/inline"
)

type BotAPI struct {
	API            *bot.Bot
	log            logger.Logger
	actionsHandler *ActionHandler
	ctx            context.Context
	cancel         context.CancelFunc
	token          string
}

func NewBot(inst *slog.Logger, token string) *BotAPI {
	return &BotAPI{
		token:          token,
		actionsHandler: NewActionHandler(),
		log:            logger.Logger{Inst: inst, Name: "BotAPI"},
	}
}

func (b *BotAPI) Init(proxy string) (err error) {
	b.log.Log(logger.INFO, "Initializing BotAPI")

	opts := []bot.Option{
		bot.WithMiddlewares(middleware.LoggingMiddleware(&b.log)),
		bot.WithDefaultHandler(func(ctx context.Context, bot *bot.Bot, update *models.Update) {
			b.log.Log(logger.INFO, "Default handler", "message", update.Message.Text)
		}),
	}

	if proxy != "" {
		proxyURL, _ := url.Parse(proxy)
		transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}

		httpClient := &http.Client{
			Transport: transport,
			Timeout:   time.Minute,
		}
		opts = append(opts, bot.WithHTTPClient(time.Minute, httpClient))
	}

	b.API, err = bot.New(b.token, opts...)
	return err
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

func (b *BotAPI) RegisterActionHandler(actionType domain.ActionType, action Action) {
	b.actionsHandler.RegisterHandler(actionType, action)
}

func (b *BotAPI) SendMessage(params *bot.SendMessageParams) error {
	_, err := b.API.SendMessage(b.ctx, params)
	if err != nil {
		b.LoggingErrors(err)
	}
	return err
}

func (b *BotAPI) SendVideoNote(params *bot.SendVideoNoteParams) error {
	_, err := b.API.SendVideoNote(b.ctx, params)
	if err != nil {
		b.LoggingErrors(err)
	}
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
