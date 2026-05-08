package tg_bot_service

import (
	"fmt"
	"sort"
	"strings"
)

// SetMaintainers parses "<service> <member1>,<member2>,..." (whitespace splits service from list)
// and updates the maintainer map. Empty member list removes the entry.
func (s *Service) SetMaintainers(raw string) string {
	service, rest := splitServiceAndRest(raw)
	if service == "" {
		return "Использование: /set_maintainers <service> <@user1>,<@user2>,..."
	}

	members := s.messagesStorage.SetMaintainers(service, rest)
	if len(members) == 0 {
		return fmt.Sprintf("Мейнтейнеры для %q очищены", service)
	}
	return fmt.Sprintf("Мейнтейнеры для %q: %s", service, strings.Join(members, ", "))
}

func (s *Service) DelMaintainers(raw string) string {
	service := strings.TrimSpace(raw)
	if service == "" {
		return "Использование: /del_maintainers <service>"
	}
	if s.messagesStorage.DeleteMaintainers(service) {
		return fmt.Sprintf("Мейнтейнеры для %q удалены", service)
	}
	return fmt.Sprintf("Сервис %q не найден", service)
}

func (s *Service) ListMaintainers() string {
	all := s.messagesStorage.ListMaintainers()
	if len(all) == 0 {
		return "Мейнтейнеры не настроены"
	}

	keys := make([]string, 0, len(all))
	for k := range all {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var b strings.Builder
	b.WriteString("Мейнтейнеры по сервисам:\n")
	for _, k := range keys {
		fmt.Fprintf(&b, "• %s — %s\n", k, strings.Join(all[k], ", "))
	}
	return b.String()
}

func splitServiceAndRest(raw string) (string, string) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", ""
	}
	idx := strings.IndexAny(raw, " \t\n")
	if idx == -1 {
		return raw, ""
	}
	return raw[:idx], strings.TrimSpace(raw[idx+1:])
}
