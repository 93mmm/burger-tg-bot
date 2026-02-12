package messages

import (
	"math/rand/v2"
	"strings"

	"github.com/93mmm/burger-tg-bot.git/internal/domain/definitions"
	"github.com/pkg/errors"
)

func (s *Storage) GetContainsMessage(key string) (definitions.Message, error) {
	var msg definitions.Message

	switch {
	case strings.Contains(key, "обед ") || strings.Contains(key, " обед"):
		msg = &definitions.TextMessage{
			Text: "Приятного аппетита",
		}
	case strings.Contains(key, "бургер"):
		msg = &definitions.TextMessage{
			Text: "господи закажи меня прямо сейчас",
		}
	case strings.Contains(key, s.gitMrURL):
		msg = &definitions.TextMessage{
			Text: "накидайте аппрувов люто\n" + s.pickRandomMembers(2),
		}
	case strings.Contains(key, "дейли"):
		msg = &definitions.TextMessage{
			Text: s.dailyLink,
		}
	case strings.Contains(key, "дэйли"):
		msg = &definitions.TextMessage{
			Text: s.dailyLink,
		}
	case strings.Contains(key, "ДИТ"):
		msg = &definitions.Gif{
			FileID: s.ditGifID,
			Quote:  "ДИТ",
		}
	default:
		return nil, errors.Wrap(definitions.ErrNotFound, "не нашли что отправлять")
	}

	return msg.Clone(), nil
}

func (s *Storage) pickRandomMembers(n int) string {
	members := s.GetMembers()
	if len(members) == 0 {
		return ""
	}
	if n > len(members) {
		n = len(members)
	}

	perm := rand.Perm(len(members))
	picked := make([]string, n)
	for i := 0; i < n; i++ {
		picked[i] = members[perm[i]]
	}

	return strings.Join(picked, " ")
}
