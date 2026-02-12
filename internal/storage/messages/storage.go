package messages

import (
	"strings"
	"sync"
)

type Storage struct {
	dailyLink string
	gitMrURL  string
	ditGifID  string

	mu      sync.RWMutex
	members []string
}

func NewStorage(dailyLink, gitMrURL, ditGifID, groupMembers string) *Storage {
	return &Storage{
		dailyLink: dailyLink,
		gitMrURL:  gitMrURL,
		ditGifID:  ditGifID,
		members:   parseMembers(groupMembers),
	}
}

func (s *Storage) SetMembers(raw string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.members = parseMembers(raw)
}

func (s *Storage) GetMembers() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]string, len(s.members))
	copy(result, s.members)
	return result
}

func parseMembers(raw string) []string {
	var members []string
	for m := range strings.SplitSeq(raw, ",") {
		m = strings.TrimSpace(m)
		if m != "" {
			members = append(members, m)
		}
	}
	return members
}
