package tvmaze

import (
	"testing"
	"time"
)

func TestCacheServesFreshEntriesAndDropsStaleOnes(t *testing.T) {
	now := time.Date(2026, 11, 16, 12, 0, 0, 0, time.UTC)
	store := newCacheStore(func() time.Time { return now })

	store.set("/search/shows?q=girls", []byte("girls"))

	now = now.Add(cacheTTL - time.Second)
	if body, ok := store.get("/search/shows?q=girls"); !ok || string(body) != "girls" {
		t.Fatalf("до истечения 5 минут ответ берётся из кеша, получили %q %v", body, ok)
	}

	now = now.Add(time.Second)
	if _, ok := store.get("/search/shows?q=girls"); ok {
		t.Fatal("через 5 минут кеш устаревает")
	}

	store.set("/shows/1", []byte("show"))
	if len(store.entries) != 1 {
		t.Fatalf("при записи устаревшие записи удаляются, осталось %d", len(store.entries))
	}
}
