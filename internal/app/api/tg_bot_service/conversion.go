package tg_bot_service

import (
	"github.com/93mmm/burger-tg-bot.git/internal/domain/definitions"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func internalMessageToExternal(msg *definitions.TextMessage) *bot.SendMessageParams {
	keyboard := [][]models.InlineKeyboardButton{}
	for _, buttonRow := range msg.Buttons {
		row := []models.InlineKeyboardButton{}
		for _, button := range buttonRow {
			btn := models.InlineKeyboardButton{
				Text:         button.Text,
				CallbackData: button.Data,
				URL:          button.URL,
			}
			row = append(row, btn)
		}
		keyboard = append(keyboard, row)
	}

	return &bot.SendMessageParams{
		ReplyParameters: &models.ReplyParameters{
			MessageID:                msg.ReplyMessageID,
			ChatID:                   msg.ChatID,
			AllowSendingWithoutReply: true,
		},
		ChatID:    msg.ChatID,
		Text:      msg.Text,
		ParseMode: models.ParseModeHTML,
		LinkPreviewOptions: &models.LinkPreviewOptions{
			IsDisabled: bot.True(),
		},
		ReplyMarkup: &models.InlineKeyboardMarkup{
			InlineKeyboard: keyboard,
		},
	}
}
