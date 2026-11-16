package shows

import (
	"context"
	"strings"

	"github.com/wildbober228/it-sinyak-homework_1_backend/tvmaze"
)

func searchShows(ctx context.Context, query searchQuery) (showsPage, error) {
	text := strings.ToLower(strings.TrimSpace(query.Text))

	if text == "" {
		return paginate(nil, query.Genre, query.Page), nil
	}

	found, err := tvmaze.SearchShows(ctx, text)
	if err != nil {
		return showsPage{}, err
	}

	return paginate(found, query.Genre, query.Page), nil
}
