package shows

import (
	"slices"
	"testing"

	"github.com/wildbober228/it-sinyak-homework_1_backend/tvmaze"
)

func catalog() []tvmaze.Show {
	genres := [][]string{
		{"Drama"}, {"Comedy", "Drama"}, {"Drama"}, {"Horror"}, {"Drama"}, {"Drama", "Horror"},
	}
	shows := make([]tvmaze.Show, 0, len(genres))

	for index, list := range genres {
		shows = append(shows, tvmaze.Show{ID: index + 1, Genres: list})
	}

	return shows
}

func TestPaginateFiltersByGenreAndCountsAllGenres(t *testing.T) {
	page := paginate(catalog(), "Drama", 2)

	if page.Total != 6 || page.Found != 5 || page.Page != 2 || page.PageCount != 2 {
		t.Fatalf("неверные счётчики: %+v", page)
	}

	if len(page.Items) != 1 || page.Items[0].ID != 6 {
		t.Fatalf("на второй странице «Drama» один сериал с id 6, получили %+v", page.Items)
	}

	want := []genreCount{{"Comedy", 1}, {"Drama", 5}, {"Horror", 2}}
	if !slices.Equal(page.Genres, want) {
		t.Fatalf("жанры считаются по всей выдаче: %+v", page.Genres)
	}
}

func TestPaginateClampsPage(t *testing.T) {
	if page := paginate(catalog(), "", 99); page.Page != 2 || len(page.Items) != 2 {
		t.Fatalf("страница за границей прижимается к последней: %+v", page)
	}

	empty := paginate(nil, "", 0)
	if empty.Page != 1 || empty.PageCount != 0 || empty.Items == nil || empty.Genres == nil {
		t.Fatalf("пустая выдача — первая страница и пустые списки, а не null: %+v", empty)
	}
}
