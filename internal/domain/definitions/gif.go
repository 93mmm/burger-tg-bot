package definitions

type Gif struct {
	ReplyMessageID int
	ChatID         any
	Quote          string

	FileID string
}

func (t *Gif) SetChatID(chatID any) {
	t.ChatID = chatID
}

func (t *Gif) SetReplyMessageID(messageID int) {
	t.ReplyMessageID = messageID
}

func (t *Gif) Clone() Message {
	clone := *t
	return &clone
}
