package messages

import "github.com/93mmm/burger-tg-bot.git/internal/domain/definitions"

func (s *Storage) GetEqualsMessage(key string) (definitions.Message, error) {
	msg, ok := equalsMessages[key]
	if !ok {
		return nil, definitions.ErrNotFound
	}

	return msg.Clone(), nil
}
