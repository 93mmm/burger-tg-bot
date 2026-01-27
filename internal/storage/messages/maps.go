package messages

import "github.com/93mmm/burger-tg-bot.git/internal/domain/definitions"

var (
	equalsMessages = map[string]definitions.Message{
		"господизаберименяотсюда": &definitions.TextMessage{
			Text: "Хлопчик, остуди свою жепу, услышал тебя родной",
		},
		"видишьfrenchbakery": &definitions.TextMessage{
			Text: "А как на счет FUCK PARENTS?",
		},
		"продупал": &definitions.TextMessage{
			Text: "Я пока не у компа, не смогу быстро посмотреть",
		},
		"тесткомандаупала": &definitions.TextMessage{
			Text: "Я не у компа сейчас, не смогу оперативно посмотреть",
		},
		"накидайте": &definitions.TextMessage{
			Text: "Накидайте за обе щеки бедолаге, благословлю",
		},
	}

	containsMessages = map[string]definitions.Message{
		"обед": &definitions.TextMessage{
			Text: "Приятного аппетита",
		},
		"бургер": &definitions.TextMessage{
			Text: "господи закажи меня прямо сейчас",
		},
		"https://git.mos.ru/buch-cloud/moscow-team-2.0": &definitions.TextMessage{
			Text: "накидайте ему аппрувов",
		},
		"дейли": &definitions.TextMessage{
			Text: "https://telemost.yandex.ru/j/24504696564321",
		},
		"дэйли": &definitions.TextMessage{
			Text: "https://telemost.yandex.ru/j/24504696564321",
		},
		"ДИТ": &definitions.Gif{
			FileID: "CgACAgIAAyEFAATnxGR_AAM5aXjkjlnvPpTK-KJTv2f9bAl5pboAAlldAALkWKBKPDih-d5XBL04BA",
			Quote:  "ДИТ",
		},
	}
)
