package definitions

type Message interface {
	SetChatID(chatID any)
	SetReplyMessageID(messageID int)
	Clone() Message
}
