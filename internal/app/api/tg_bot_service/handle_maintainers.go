package tg_bot_service

import (
	"context"
	"strings"

	"github.com/93mmm/burger-tg-bot.git/internal/utils/logger"
	"github.com/93mmm/burger-tg-bot.git/internal/utils/tgbot"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (s *Server) HandleSetMaintainers(ctx context.Context, b *bot.Bot, update *models.Update) {
	if !s.isAdminInPrivate(update) {
		return
	}

	raw := strings.TrimSpace(strings.TrimPrefix(update.Message.Text, "/set_maintainers"))
	s.replyToAdmin(ctx, b, update, s.service.SetMaintainers(raw))
}

func (s *Server) HandleDelMaintainers(ctx context.Context, b *bot.Bot, update *models.Update) {
	if !s.isAdminInPrivate(update) {
		return
	}

	raw := strings.TrimSpace(strings.TrimPrefix(update.Message.Text, "/del_maintainers"))
	s.replyToAdmin(ctx, b, update, s.service.DelMaintainers(raw))
}

func (s *Server) HandleListMaintainers(ctx context.Context, b *bot.Bot, update *models.Update) {
	if !s.isAdminInPrivate(update) {
		return
	}

	s.replyToAdmin(ctx, b, update, s.service.ListMaintainers())
}

func (s *Server) replyToAdmin(ctx context.Context, b *bot.Bot, update *models.Update, text string) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: tgbot.GetChatID(update),
		Text:   text,
	})
	if err != nil {
		logger.ErrorKV(ctx, "error sending admin response", err)
	}
}
