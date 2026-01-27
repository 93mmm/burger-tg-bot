package tg_bot_service

import (
	"strings"

	"github.com/93mmm/burger-tg-bot.git/internal/services/tg_bot_service/clients"
)

type Service struct {
	replacer        *strings.Replacer
	messagesStorage clients.MessagesStorage
}

func NewService() *Service {
	replacer := strings.NewReplacer(
		",", "",
		".", "",
		" ", "",
		"?", "",
		"!", "",
		")", "",
		"(", "",
		"`", "",
	)

	return &Service{
		replacer: replacer,
	}
}
