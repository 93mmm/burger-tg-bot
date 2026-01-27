package tgbot

import (
	"github.com/93mmm/burger-tg-bot.git/internal/domain/definitions"
	"github.com/go-telegram/bot/models"
)

func GetUserID(update *models.Update) definitions.TelegramUserID {
	if update.Message != nil {
		return update.Message.From.ID
	}

	if update.CallbackQuery != nil {
		return update.CallbackQuery.From.ID
	}

	return 0
}
