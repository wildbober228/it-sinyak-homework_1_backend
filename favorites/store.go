package favorites

import (
	"maps"
	"slices"
	"sync"

	"github.com/wildbober228/it-sinyak-homework_1_backend/tvmaze"
)

type favoritesStore struct {
	mu      sync.Mutex
	clients map[string]map[int]favorite
	added   int
}

var store = newFavoritesStore()

func newFavoritesStore() *favoritesStore {
	return &favoritesStore{clients: map[string]map[int]favorite{}}
}

func (s *favoritesStore) has(clientID string, showID int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.clients[clientID][showID]

	return ok
}

func (s *favoritesStore) add(clientID string, show tvmaze.Show) {
	s.mu.Lock()
	defer s.mu.Unlock()

	shows, ok := s.clients[clientID]
	if !ok {
		shows = map[int]favorite{}
		s.clients[clientID] = shows
	}

	if _, exists := shows[show.ID]; exists {
		return
	}

	s.added++
	shows[show.ID] = favorite{show: show, order: s.added}
}

func (s *favoritesStore) remove(clientID string, showIDs []int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	shows := s.clients[clientID]

	for _, id := range showIDs {
		delete(shows, id)
	}

	if len(shows) == 0 {
		delete(s.clients, clientID)
	}
}

func (s *favoritesStore) list(clientID string) []tvmaze.Show {
	s.mu.Lock()
	defer s.mu.Unlock()

	entries := slices.Collect(maps.Values(s.clients[clientID]))

	slices.SortFunc(entries, func(a, b favorite) int {
		return b.order - a.order
	})

	shows := make([]tvmaze.Show, 0, len(entries))

	for _, entry := range entries {
		shows = append(shows, entry.show)
	}

	return shows
}
