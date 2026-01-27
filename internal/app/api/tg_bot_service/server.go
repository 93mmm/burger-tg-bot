package tg_bot_service

import (
	"context"

	"github.com/93mmm/burger-tg-bot.git/internal/app/api/middlewares"
	"github.com/93mmm/burger-tg-bot.git/internal/domain/definitions"
	"github.com/93mmm/burger-tg-bot.git/internal/utils/logger"
	"github.com/go-telegram/bot"
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
	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix, s.HandleAnyMessage, middlewares.PanicRecoveryMiddleware(), middlewares.InoutLogging())
}

func (s *Server) Start(ctx context.Context, b *bot.Bot) {
	b.Start(ctx)
}

// Sends message
func (s *Server) sendMessage(ctx context.Context, b *bot.Bot, msg *definitions.Message) {
	if msg == nil {
		logger.ErrorKV(ctx, "message is nil!")
		return
	}
	tgMessage := internalMessageToExternal(msg)

	logger.DebugKV(ctx, "method called", "m", tgMessage)

	_, err := b.SendMessage(
		ctx,
		tgMessage,
	)
	if err != nil {
		logger.ErrorKV(ctx, "error happened while sending message", err)
	}
}
