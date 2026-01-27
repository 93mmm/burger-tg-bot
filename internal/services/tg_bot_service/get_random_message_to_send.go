package tg_bot_service

import (
	"context"
	"math/rand/v2"
	"strings"

	"github.com/93mmm/burger-tg-bot.git/internal/domain/definitions"
	"github.com/93mmm/burger-tg-bot.git/internal/utils/logger"
)

var (
	randomMessages = map[string]*definitions.Message{
		"заберименяпрямосейчас": &definitions.Message{},
	}
)

func (s *Service) GetRandomMessageToSend(ctx context.Context, text string, chatID any) (*definitions.Message, error) {
	key := s.textToNormalForm(text)

	message, ok := randomMessages[text]
	if !ok {
		logger.InfoKV(ctx, "не отправили потому что нет такого ключа", "key", key)
		return nil, nil
	}

	if toSendOrNotToSend() {
		return message, nil
	}

	logger.InfoKV(ctx, "не отправили потому что toSendOrNotToSend решил не отправлять", "key", key)
	return nil, nil
}

func (s *Service) textToNormalForm(text string) string {
	text = s.replacer.Replace(text)
	return strings.ToLower(text)
}

func toSendOrNotToSend() bool {
	return rand.IntN(10) > 5
}
