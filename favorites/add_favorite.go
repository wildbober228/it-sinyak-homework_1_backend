package favorites

import (
	"context"

	"github.com/wildbober228/it-sinyak-homework_1_backend/tvmaze"
)

func addFavorite(ctx context.Context, clientID string, showID int) ([]tvmaze.Show, error) {
	if !store.has(clientID, showID) {
		show, err := tvmaze.GetShow(ctx, showID)
		if err != nil {
			return nil, err
		}

		store.add(clientID, show)
	}

	return store.list(clientID), nil
}
