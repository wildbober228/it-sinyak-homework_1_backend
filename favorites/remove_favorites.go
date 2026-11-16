package favorites

import "github.com/wildbober228/it-sinyak-homework_1_backend/tvmaze"

func removeFavorites(clientID string, showIDs []int) []tvmaze.Show {
	store.remove(clientID, showIDs)

	return store.list(clientID)
}
