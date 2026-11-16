package favorites

import "github.com/wildbober228/it-sinyak-homework_1_backend/tvmaze"

type favorite struct {
	show  tvmaze.Show
	order int
}

type favoritesList struct {
	Items []tvmaze.Show `json:"items"`
}
