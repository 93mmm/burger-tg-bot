package messages

import (
	"context"
	"math/rand/v2"
	"regexp"
	"strings"
	"time"

	"github.com/93mmm/burger-tg-bot.git/internal/domain/definitions"
	"github.com/93mmm/burger-tg-bot.git/internal/utils/logger"
	"github.com/pkg/errors"
)

var moscowTZ = time.FixedZone("MSK", 3*60*60)

// Captures the project segment right before /-/merge_requests/N in a GitLab URL.
var mrServiceRe = regexp.MustCompile(`/([^/\s]+/[^/\s]+)/-/merge_requests/\d+`)

func (s *Storage) GetContainsMessage(key string) (definitions.Message, error) {
	var msg definitions.Message

	switch {
	case strings.Contains(key, "обед ") || strings.Contains(key, " обед"):
		msg = &definitions.TextMessage{
			Text: "Приятного аппетита",
		}
	case strings.Contains(key, "бургер"):
		msg = &definitions.TextMessage{
			Text: "господи закажи меня прямо сейчас",
		}
	case strings.Contains(key, s.gitMrURL):
		text := "Накидайте аппрувов"
		if isWorkingHours() {
			if tags := s.tagsForMR(key); tags != "" {
				text += "\n\n" + tags
			}
		}
		msg = &definitions.TextMessage{
			Text: text,
		}
	case strings.Contains(key, "дейли"):
		msg = &definitions.TextMessage{
			Text: s.dailyLink,
		}
	case strings.Contains(key, "дэйли"):
		msg = &definitions.TextMessage{
			Text: s.dailyLink,
		}
	case strings.Contains(key, "ДИТ"):
		msg = &definitions.Gif{
			FilePath: DitGifPath,
			Quote:    "ДИТ",
		}
	default:
		return nil, errors.Wrap(definitions.ErrNotFound, "не нашли что отправлять")
	}

	return msg.Clone(), nil
}

// tagsForMR builds the mention block for an MR message:
//   - service has ≥2 maintainers → "Мейнтейнеры: <all of them>"
//   - service has exactly 1 maintainer → "Мейнтейнеры: <maintainer> <random buddy>"
//     (random pulled from default pool, excluding the maintainer to avoid duplicates)
//   - service unknown / no maintainers → 2 random members from default pool, no label
//
// Returns "" if there's nobody to tag at all.
func (s *Storage) tagsForMR(text string) string {
	var maintainers []string
	service := extractServiceName(text)
	if service != "" {
		maintainers = s.GetMaintainers(service)
	}

	logger.DebugKV(context.Background(), "building tags for MR", "service", service, "maintainers", maintainers)

	if len(maintainers) == 0 {
		random := s.pickRandomMembers(2, nil)
		if len(random) == 0 {
			return ""
		}
		return strings.Join(random, " ")
	}

	maintanerTags := append([]string{}, maintainers...)
	var randomBuddy []string
	if len(maintainers) == 1 {
		randomBuddy = s.pickRandomMembers(1, maintainers)
		if len(randomBuddy) != 0 {
			return "<b>Мейнтейнеры:</b> " + strings.Join(maintanerTags, " ") + "\n\nИ пж ещё апрувы от " + randomBuddy[0]
		}
	}
	return "<b>Мейнтейнеры:</b> " + strings.Join(maintanerTags, " ")
}

func extractServiceName(text string) string {
	m := mrServiceRe.FindStringSubmatch(text)
	if len(m) >= 2 {
		return m[1]
	}
	return ""
}

// pickRandomMembers picks up to n distinct members from the default pool,
// skipping anyone listed in `exclude` so MR mentions never tag the same person twice.
func (s *Storage) pickRandomMembers(n int, exclude []string) []string {
	if n <= 0 {
		return nil
	}
	members := s.GetMembers()
	if len(members) == 0 {
		return nil
	}

	excludeSet := make(map[string]struct{}, len(exclude))
	for _, e := range exclude {
		excludeSet[e] = struct{}{}
	}

	pool := make([]string, 0, len(members))
	for _, m := range members {
		if _, skip := excludeSet[m]; !skip {
			pool = append(pool, m)
		}
	}
	if len(pool) == 0 {
		return nil
	}
	if n > len(pool) {
		n = len(pool)
	}

	perm := rand.Perm(len(pool))
	picked := make([]string, n)
	for i := 0; i < n; i++ {
		picked[i] = pool[perm[i]]
	}
	return picked
}

func isWorkingHours() bool {
	now := time.Now().In(moscowTZ)
	weekday := now.Weekday()
	hour := now.Hour()
	return weekday >= time.Monday && weekday <= time.Friday && hour >= 9 && hour < 18 || true
}
