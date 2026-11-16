package favorites

import "github.com/wildbober228/it-sinyak-homework_1_backend/tvmaze"

func listFavorites(clientID string) []tvmaze.Show {
	return store.list(clientID)
}
