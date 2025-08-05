package cache

import (
	"sync"
	"time"
)

type DuplicateStore struct {
	mu    sync.Mutex
	items map[string]time.Time
	ttl   time.Duration
}

func NewDuplicateStore(ttl time.Duration) *DuplicateStore {
	return &DuplicateStore{
		items: make(map[string]time.Time),
		ttl:   ttl,
	}
}

func (s *DuplicateStore) Seen(fingerprint string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	expiry, found := s.items[fingerprint]
	if found && time.Now().Before(expiry) {
		return true
	}

	s.items[fingerprint] = time.Now().Add(s.ttl)
	return false
}
