package tg_bot_service

import (
	"context"

	"github.com/93mmm/burger-tg-bot.git/internal/domain/definitions"
	"github.com/93mmm/burger-tg-bot.git/internal/utils/logger"
	"github.com/93mmm/burger-tg-bot.git/internal/utils/tgbot"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/pkg/errors"
)

func (s *Server) HandleAnyMessage(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		logger.Warn(ctx, "update.Message is nil")
		return
	}

	messageText := update.Message.Text
	chatID := tgbot.GetChatID(update)
	messageID := update.Message.ID

	message, err := s.service.GetRandomMessageToSend(ctx, messageText, chatID, messageID)
	if err != nil {
		if errors.Is(err, definitions.ErrNotFound) || errors.Is(err, definitions.ErrDecidedToNotSend) {
			return
		}
		logger.ErrorKV(ctx, "unexpected error happened", err)
		return
	}

	s.sendMessage(ctx, b, message)
}
