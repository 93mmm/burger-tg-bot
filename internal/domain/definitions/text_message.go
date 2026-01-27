package definitions

type Button struct {
	Text string
	Data string
	URL  string
}

type TextMessage struct {
	ReplyMessageID int
	ChatID         any
	Text           string

	Buttons [][]Button
}

func (m *TextMessage) SetChatID(chatID any) {
	m.ChatID = chatID
}

func (m *TextMessage) SetReplyMessageID(messageID int) {
	m.ReplyMessageID = messageID
}

func (m *TextMessage) Clone() Message {
	clone := *m

	if m.Buttons != nil {
		newButtons := make([][]Button, len(m.Buttons))
		for i := range m.Buttons {
			if m.Buttons[i] != nil {
				row := make([]Button, len(m.Buttons[i]))
				copy(row, m.Buttons[i])
				newButtons[i] = row
			}
		}
		clone.Buttons = newButtons
	}

	return &clone
}
