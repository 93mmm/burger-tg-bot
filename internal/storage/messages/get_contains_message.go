package messages

import (
	"strings"

	"github.com/93mmm/burger-tg-bot.git/internal/domain/definitions"
	"github.com/pkg/errors"
)

func (s *Storage) GetContainsMessage(key string) (definitions.Message, error) {
	for substr, msg := range containsMessages {
		if strings.Contains(key, substr) {
			return msg.Clone(), nil
		}
	}

	return nil, errors.Wrap(definitions.ErrNotFound, "не нашли что отправлять")
}
