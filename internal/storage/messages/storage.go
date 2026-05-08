package messages

import (
	"strings"
	"sync"
)

// DitGifPath is the on-disk location of the ДИТ animation, copied into the
// Docker image from ./assets. Always re-uploaded via multipart so we never
// hit the "file_id went stale" issue.
const DitGifPath = "assets/god.mp4"

type Storage struct {
	dailyLink string
	gitMrURL  string

	mu          sync.RWMutex
	members     []string
	maintainers map[string][]string
}

func NewStorage(dailyLink, gitMrURL, groupMembers string) (*Storage, error) {
	s := &Storage{
		dailyLink:   dailyLink,
		gitMrURL:    gitMrURL,
		members:     parseMembers(groupMembers),
		maintainers: make(map[string][]string),
	}
	if err := s.loadMaintainers(); err != nil {
		return nil, err
	}
	return s, nil
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

// SetMaintainers replaces the maintainer list for the given service and persists.
// Empty raw removes the entry. Returns the parsed members for confirmation.
func (s *Storage) SetMaintainers(service, raw string) []string {
	members := parseMembers(raw)
	s.mu.Lock()
	if len(members) == 0 {
		delete(s.maintainers, service)
	} else {
		s.maintainers[service] = members
	}
	s.mu.Unlock()

	s.saveMaintainers()

	if len(members) == 0 {
		return nil
	}
	cp := make([]string, len(members))
	copy(cp, members)
	return cp
}

func (s *Storage) DeleteMaintainers(service string) bool {
	s.mu.Lock()
	_, ok := s.maintainers[service]
	if ok {
		delete(s.maintainers, service)
	}
	s.mu.Unlock()

	if !ok {
		return false
	}
	s.saveMaintainers()
	return true
}

func (s *Storage) GetMaintainers(service string) []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	mts, ok := s.maintainers[service]
	if !ok {
		return nil
	}
	result := make([]string, len(mts))
	copy(result, mts)
	return result
}

func (s *Storage) ListMaintainers() map[string][]string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make(map[string][]string, len(s.maintainers))
	for k, v := range s.maintainers {
		cp := make([]string, len(v))
		copy(cp, v)
		result[k] = cp
	}
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
