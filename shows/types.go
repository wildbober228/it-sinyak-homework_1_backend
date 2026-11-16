package shows

import "github.com/wildbober228/it-sinyak-homework_1_backend/tvmaze"

type searchQuery struct {
	Text  string
	Genre string
	Page  int
}

type genreCount struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type showsPage struct {
	Items     []tvmaze.Show `json:"items"`
	Genres    []genreCount  `json:"genres"`
	Total     int           `json:"total"`
	Found     int           `json:"found"`
	Page      int           `json:"page"`
	PageCount int           `json:"pageCount"`
}
