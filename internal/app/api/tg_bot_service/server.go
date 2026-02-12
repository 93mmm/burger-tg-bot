package tg_bot_service

import (
	"context"

	"github.com/93mmm/burger-tg-bot.git/internal/app/api/middlewares"
	"github.com/93mmm/burger-tg-bot.git/internal/domain/definitions"
	"github.com/93mmm/burger-tg-bot.git/internal/utils/debug"
	"github.com/93mmm/burger-tg-bot.git/internal/utils/logger"
	"github.com/go-telegram/bot"
	"github.com/pkg/errors"
)

// Transport layer in telegram bot
type Server struct {
	service Service
}

func NewServer(
	service Service,
) *Server {
	srv := &Server{
		service: service,
	}
	return srv
}

func (s *Server) RegisterHandlers(b *bot.Bot) {
	b.RegisterHandler(bot.HandlerTypeMessageText, "/change_tags", bot.MatchTypePrefix, s.HandleChangeTags, middlewares.PanicRecoveryMiddleware(), middlewares.InoutLogging())
	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix, s.HandleAnyMessage, middlewares.PanicRecoveryMiddleware(), middlewares.InoutLogging())
}

func (s *Server) Start(ctx context.Context, b *bot.Bot) {
	b.Start(ctx)
}

// Sends message
func (s *Server) sendMessage(ctx context.Context, b *bot.Bot, msg definitions.Message) {
	if msg == nil {
		logger.ErrorKV(ctx, "message is nil!")
		return
	}
	logger.DebugKV(ctx, "trying to send message", "message", msg)
	var err error

	switch m := msg.(type) {
	case *definitions.TextMessage:
		tgMessage := internalMessageToExternal(m)
		_, err = b.SendMessage(
			ctx,
			tgMessage,
		)

	case *definitions.Gif:
		tgMessage := internalGifToExternal(m)
		_, err = b.SendAnimation(
			ctx,
			tgMessage,
		)

	default:
		err = errors.Wrapf(definitions.ErrInternal, "received undefined type of message: %v", debug.TypeOf(msg))
	}

	if err != nil {
		logger.ErrorKV(ctx, "error happened while sending message", err)
	}
}
