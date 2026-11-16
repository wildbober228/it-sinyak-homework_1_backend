package tvmaze

import "time"

type Show struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	URL     string   `json:"url"`
	Image   *string  `json:"image"`
	Genres  []string `json:"genres"`
	Year    *string  `json:"year"`
	Rating  *string  `json:"rating"`
	Network *string  `json:"network"`
	Summary string   `json:"summary"`
}

type apiShow struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	URL        string      `json:"url"`
	Genres     []string    `json:"genres"`
	Premiered  *string     `json:"premiered"`
	Rating     apiRating   `json:"rating"`
	Network    *apiChannel `json:"network"`
	WebChannel *apiChannel `json:"webChannel"`
	Image      *apiImage   `json:"image"`
	Summary    *string     `json:"summary"`
}

type apiRating struct {
	Average *float64 `json:"average"`
}

type apiChannel struct {
	Name string `json:"name"`
}

type apiImage struct {
	Medium string `json:"medium"`
}

type apiSearchResult struct {
	Show apiShow `json:"show"`
}

type cacheEntry struct {
	body     []byte
	storedAt time.Time
}
