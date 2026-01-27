package tg_bot_service

import "strings"

type Service struct {
	replacer *strings.Replacer
}

func NewService() *Service {
	replacer := strings.NewReplacer(
		",", "",
		".", "",
		" ", "",
		"?", "",
		"!", "",
		")", "",
		"(", "",
		"`", "",
	)

	return &Service{
		replacer: replacer,
	}
}
