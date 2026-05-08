package clients

import "github.com/93mmm/burger-tg-bot.git/internal/domain/definitions"

type MessagesStorage interface {
	GetEqualsMessage(key string) (definitions.Message, error)
	GetContainsMessage(key string) (definitions.Message, error)
	SetMembers(raw string)
	GetMembers() []string
	SetMaintainers(service, raw string) []string
	DeleteMaintainers(service string) bool
	GetMaintainers(service string) []string
	ListMaintainers() map[string][]string
}
