package tg_bot_service

import (
	"context"
	"strings"

	"github.com/93mmm/burger-tg-bot.git/internal/utils/logger"
	"github.com/93mmm/burger-tg-bot.git/internal/utils/tgbot"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (s *Server) HandleChangeTags(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	// работает только в личных сообщениях
	if update.Message.Chat.Type != "private" {
		return
	}

	raw := strings.TrimPrefix(update.Message.Text, "/change_tags")
	raw = strings.TrimSpace(raw)

	response := s.service.ChangeTags(raw)

	chatID := tgbot.GetChatID(update)
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   response,
	})
	if err != nil {
		logger.ErrorKV(ctx, "error sending change_tags response", err)
	}
}
