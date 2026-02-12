package tg_bot_service

import "strings"

func (s *Service) ChangeTags(raw string) string {
	s.messagesStorage.SetMembers(raw)

	members := s.messagesStorage.GetMembers()
	if len(members) == 0 {
		return "Список тегов очищен"
	}
	return "Теги обновлены: " + strings.Join(members, ", ")
}
