package messages

import "github.com/93mmm/burger-tg-bot.git/internal/domain/definitions"

func (s *Storage) GetEqualsMessage(key string) (definitions.Message, error) {
	var msg definitions.Message

	switch key {
	case "господизаберименяотсюда":
		msg = &definitions.TextMessage{
			Text: "Хлопчик, остуди свою жепу, услышал тебя родной",
		}
	case "видишьfrenchbakery":
		msg = &definitions.TextMessage{
			Text: "А как на счет FUCK PARENTS?",
		}
	case "продупал":
		msg = &definitions.TextMessage{
			Text: "Я пока не у компа, не смогу быстро посмотреть",
		}
	case "тесткомандаупала":
		msg = &definitions.TextMessage{
			Text: "Я не у компа сейчас, не смогу оперативно посмотреть",
		}
	case "накидайте":
		msg = &definitions.TextMessage{
			Text: "Накидайте за обе щеки бедолаге, благословлю",
		}
	default:
		return nil, definitions.ErrNotFound
	}

	return msg.Clone(), nil
}
