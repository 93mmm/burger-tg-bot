package tg_bot_service

import (
	"context"

	"github.com/93mmm/burger-tg-bot.git/internal/utils/logger"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (s *Server) HandleAnyMessage(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		logger.Warn(ctx, "update.Message is nil")
		return
	}

	messageText := update.Message.Text

	message, err := s.service.GetRandomMessageToSend(ctx, messageText, update.Message.Chat.ID)
	if err != nil {
		return
	}

	s.sendMessage(ctx, b, message)
}
