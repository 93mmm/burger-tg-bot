package tg_bot_service

import (
	"context"
	"math/rand/v2"
	"strings"

	"github.com/93mmm/burger-tg-bot.git/internal/domain/definitions"
	"github.com/pkg/errors"
)

func (s *Service) GetRandomMessageToSend(ctx context.Context, text string, chatID any, messageID int) (definitions.Message, error) {
	message, err := s.messagesStorage.GetEqualsMessage(
		s.textToNormalForm(text),
	)
	if err == nil {
		// рандомно выбираем будем ли отправлять сообщение
		if rand.IntN(10) > 5 {
			message.SetChatID(chatID)
			return message, nil
		}
		return nil, definitions.ErrDecidedToNotSend
	}
	if !errors.Is(err, definitions.ErrDecidedToNotSend) && !errors.Is(err, definitions.ErrNotFound) {
		return nil, errors.Wrap(err, "произошла неизвестная ошибка")
	}

	message, err = s.messagesStorage.GetContainsMessage(text)
	if err == nil {
		message.SetChatID(chatID)
		message.SetReplyMessageID(messageID)
		return message, nil
	}
	if !errors.Is(err, definitions.ErrDecidedToNotSend) && !errors.Is(err, definitions.ErrNotFound) {
		return nil, errors.Wrap(err, "произошла неизвестная ошибка")
	}

	return nil, errors.Wrap(definitions.ErrNotFound, "не придумали что отвечать")
}

func (s *Service) textToNormalForm(text string) string {
	text = s.replacer.Replace(text)
	return strings.ToLower(text)
}
