package shows

import (
	"slices"
	"strings"

	"github.com/wildbober228/it-sinyak-homework_1_backend/tvmaze"
)

func paginate(shows []tvmaze.Show, genre string, page int) showsPage {
	filtered := filterByGenre(shows, genre)
	pageCount := (len(filtered) + pageSize - 1) / pageSize
	current := min(max(page, 1), max(pageCount, 1))
	start := min((current-1)*pageSize, len(filtered))
	end := min(start+pageSize, len(filtered))

	return showsPage{
		Items:     append([]tvmaze.Show{}, filtered[start:end]...),
		Genres:    countGenres(shows),
		Total:     len(shows),
		Found:     len(filtered),
		Page:      current,
		PageCount: pageCount,
	}
}

func filterByGenre(shows []tvmaze.Show, genre string) []tvmaze.Show {
	if genre == "" {
		return shows
	}

	return slices.DeleteFunc(slices.Clone(shows), func(show tvmaze.Show) bool {
		return !slices.Contains(show.Genres, genre)
	})
}

func countGenres(shows []tvmaze.Show) []genreCount {
	counts := map[string]int{}

	for _, show := range shows {
		for _, genre := range show.Genres {
			counts[genre]++
		}
	}

	genres := make([]genreCount, 0, len(counts))

	for name, count := range counts {
		genres = append(genres, genreCount{Name: name, Count: count})
	}

	slices.SortFunc(genres, func(a, b genreCount) int {
		return strings.Compare(a.Name, b.Name)
	})

	return genres
}
