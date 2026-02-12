package app

import (
	"context"

	"github.com/93mmm/burger-tg-bot.git/cmd/config"
	tg_bot_server "github.com/93mmm/burger-tg-bot.git/internal/app/api/tg_bot_service"
	"github.com/93mmm/burger-tg-bot.git/internal/services/tg_bot_service"
	"github.com/93mmm/burger-tg-bot.git/internal/storage/messages"
	"github.com/93mmm/burger-tg-bot.git/internal/utils/logger"
	"github.com/go-telegram/bot"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type Server interface {
	Run(ctx context.Context) error
	Shutdown(ctx context.Context) error
}

type Environment = string

const (
	EnvironmentDevelop Environment = "develop"
	EnvironmentTest    Environment = "test"
	EnvironmentProd    Environment = "prod"
)

type app struct {
	config *config.Config
	cancel context.CancelFunc
}

func NewApp(cfg *config.Config) *app {
	return &app{
		config: cfg,
	}
}

func (a *app) Run(c context.Context) error {
	ctx, cancel := context.WithCancel(c)
	a.cancel = cancel

	context.AfterFunc(ctx, func() {
		defer func() {
			if e := recover(); e != nil {
				logger.ErrorKV(ctx, "panic: context after-func", "recover", e)
			}
		}()
		if err := a.GracefulShutdown(ctx); err != nil {
			logger.ErrorKV(ctx, "graceful shutdown error", err)
		}
	})

	opts := []bot.Option{}
	if a.config.Environment == "develop" {
		opts = append(opts, bot.WithDebug())
		logger.Debug(ctx, "bot debugging enabled")
	}

	b, err := bot.New(a.config.BotToken)
	if err != nil {
		return errors.Wrap(err, "error while creating bot instance")
	}

	storage := messages.NewStorage(
		a.config.Messages.DailyLink,
		a.config.Messages.GitMrURL,
		a.config.Messages.DitGifID,
		a.config.Messages.GroupMembers,
	)

	service := tg_bot_service.NewService(storage)

	server := tg_bot_server.NewServer(service)

	server.RegisterHandlers(b)

	eg, ctx := errgroup.WithContext(c)

	eg.Go(func() error {
		defer a.cancel()
		server.Start(ctx, b)
		return nil
	})

	logger.Info(ctx, "сервер поднялся")

	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "errorgroup вернула ошибку")
	}
	return nil
}

func (a *app) GracefulShutdown(ctx context.Context) error {
	return nil
}
