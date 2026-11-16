package tvmaze

import (
	"html"
	"regexp"
	"strconv"
	"strings"
)

var htmlTag = regexp.MustCompile(`<[^>]*>`)

func toShow(show apiShow) Show {
	return Show{
		ID:      show.ID,
		Name:    show.Name,
		URL:     show.URL,
		Image:   imageOf(show),
		Genres:  genresOf(show),
		Year:    yearOf(show.Premiered),
		Rating:  ratingOf(show.Rating.Average),
		Network: networkOf(show),
		Summary: stripHTML(show.Summary),
	}
}

func imageOf(show apiShow) *string {
	if show.Image == nil {
		return nil
	}

	return &show.Image.Medium
}

func genresOf(show apiShow) []string {
	if show.Genres == nil {
		return []string{}
	}

	return show.Genres
}

func yearOf(premiered *string) *string {
	if premiered == nil || len(*premiered) < 4 {
		return nil
	}

	year := (*premiered)[:4]

	return &year
}

func ratingOf(average *float64) *string {
	if average == nil {
		return nil
	}

	rating := strconv.FormatFloat(*average, 'f', 1, 64)

	return &rating
}

func networkOf(show apiShow) *string {
	switch {
	case show.Network != nil:
		return &show.Network.Name
	case show.WebChannel != nil:
		return &show.WebChannel.Name
	default:
		return nil
	}
}

func stripHTML(summary *string) string {
	if summary == nil {
		return ""
	}

	text := htmlTag.ReplaceAllString(*summary, "")

	return strings.TrimSpace(html.UnescapeString(text))
}
