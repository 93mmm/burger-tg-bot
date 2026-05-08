package tg_bot_service

import (
	"context"

	"github.com/93mmm/burger-tg-bot.git/internal/domain/definitions"
)

type Service interface {
	GetRandomMessageToSend(ctx context.Context, text string, chatID any, messageID int) (definitions.Message, error)
	ChangeTags(raw string) string
	SetMaintainers(raw string) string
	DelMaintainers(raw string) string
	ListMaintainers() string
}
