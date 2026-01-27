package tg_bot_service

import (
	"context"

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
	// handle commands
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, s.HandleAnyMessage)

	b.RegisterHandler(bot.HandlerTypeMessageText, "/", bot.MatchTypePrefix, s.HandleUnknownCommand)

	// case callbacks.CallbackCredits:
	// 	s.HandleCallbackCredits(ctx, update)
	// case callbacks.CallbackContactManager:
	// 	s.HandleCallbackContactManager(ctx, update)
	// case callbacks.CallbackBack:
	// 	s.HandleCallbackBack(ctx, update)

	b.RegisterHandler(bot.HandlerTypeCallbackQueryData, "btn_", bot.MatchTypePrefix, s.ButtonCallbackHandler)

	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypePrefix, s.HandleAnyText)
}

func (s *Server) Start(ctx context.Context, b *bot.Bot) {
	b.Start(ctx)
}

// Sends message
func (s *Server) sendMessage(ctx context.Context, b *bot.Bot, msg *definitions.Message) {
	logger.DebugKV(ctx, "method called", "m", msg)

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
