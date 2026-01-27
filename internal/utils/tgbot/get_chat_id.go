package tgbot

import (
	"github.com/93mmm/burger-tg-bot.git/internal/domain/definitions"
	"github.com/go-telegram/bot/models"
)

func GetChatID(update *models.Update) definitions.ChatID {
	if update.Message != nil {
		return update.Message.Chat.ID
	}

	if update.CallbackQuery != nil {
		return update.CallbackQuery.From.ID
	}

	return 0
}
