package tvmaze

import (
	"sync"
	"time"
)

type cacheStore struct {
	mu      sync.Mutex
	entries map[string]cacheEntry
	now     func() time.Time
}

var cache = newCacheStore(time.Now)

func newCacheStore(now func() time.Time) *cacheStore {
	return &cacheStore{entries: map[string]cacheEntry{}, now: now}
}

func (c *cacheStore) get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if !ok || c.isStale(entry) {
		return nil, false
	}

	return entry.body, true
}

func (c *cacheStore) set(key string, body []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for staleKey, entry := range c.entries {
		if c.isStale(entry) {
			delete(c.entries, staleKey)
		}
	}

	c.entries[key] = cacheEntry{body: body, storedAt: c.now()}
}

func (c *cacheStore) isStale(entry cacheEntry) bool {
	return c.now().Sub(entry.storedAt) >= cacheTTL
}
