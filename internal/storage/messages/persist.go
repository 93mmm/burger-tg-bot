package messages

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/93mmm/burger-tg-bot.git/internal/utils/logger"
	"github.com/pkg/errors"
)

// MaintainersFile is the on-disk JSON snapshot of the service→maintainers map.
// Relative to cwd: from project root locally, /root/data/... in the Docker image
// (where /root/data is expected to be a mounted volume).
const MaintainersFile = "data/maintainers.json"

// loadMaintainers reads MaintainersFile if it exists. Missing file is not an error —
// we treat first run as "empty map".
func (s *Storage) loadMaintainers() error {
	data, err := os.ReadFile(MaintainersFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return errors.Wrapf(err, "read %s", MaintainersFile)
	}

	var m map[string][]string
	if err := json.Unmarshal(data, &m); err != nil {
		return errors.Wrapf(err, "unmarshal %s", MaintainersFile)
	}
	if m == nil {
		m = make(map[string][]string)
	}

	s.mu.Lock()
	s.maintainers = m
	s.mu.Unlock()
	return nil
}

// saveMaintainers writes the current map atomically (tmp + rename).
// Errors are logged at error level — callers continue with in-memory state intact.
func (s *Storage) saveMaintainers() {
	s.mu.RLock()
	data, err := json.MarshalIndent(s.maintainers, "", "  ")
	s.mu.RUnlock()
	if err != nil {
		logger.ErrorKV(context.Background(), "marshal maintainers", "err", err)
		return
	}

	if err := os.MkdirAll(filepath.Dir(MaintainersFile), 0o755); err != nil {
		logger.ErrorKV(context.Background(), "mkdir maintainers data dir", "err", err)
		return
	}

	tmp := MaintainersFile + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		logger.ErrorKV(context.Background(), "write maintainers tmp", "err", err)
		return
	}
	if err := os.Rename(tmp, MaintainersFile); err != nil {
		logger.ErrorKV(context.Background(), "rename maintainers tmp", "err", err)
	}
}
