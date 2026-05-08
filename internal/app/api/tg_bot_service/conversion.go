package tg_bot_service

import (
	"io"
	"os"
	"path/filepath"

	"github.com/93mmm/burger-tg-bot.git/internal/domain/definitions"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/pkg/errors"
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

// internalGifToExternal opens the local GIF file and prepares a multipart
// upload — sending by file_id used to break randomly, multipart always works.
// Caller MUST close the returned io.Closer.
func internalGifToExternal(msg *definitions.Gif) (*bot.SendAnimationParams, io.Closer, error) {
	if msg.FilePath == "" {
		return nil, nil, errors.New("gif has empty FilePath")
	}
	f, err := os.Open(msg.FilePath)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "open gif file %q", msg.FilePath)
	}
	return &bot.SendAnimationParams{
		ReplyParameters: &models.ReplyParameters{
			MessageID:                msg.ReplyMessageID,
			ChatID:                   msg.ChatID,
			AllowSendingWithoutReply: true,
			Quote:                    msg.Quote,
		},
		ChatID: msg.ChatID,
		Animation: &models.InputFileUpload{
			Filename: filepath.Base(msg.FilePath),
			Data:     f,
		},
	}, f, nil
}
