package tvmaze

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const girlsResponse = `[{"score":1,"show":{"id":139,"name":"Girls","url":"https://www.tvmaze.com/shows/139/girls",
"genres":["Drama","Romance"],"premiered":"2012-04-15","rating":{"average":6.96},
"network":{"name":"HBO"},"webChannel":null,"image":null,"summary":"<p>Hannah &amp; <b>friends</b></p>"}}]`

func TestSearchShowsMapsResponseAndUsesCache(t *testing.T) {
	requests := useFakeTVmaze(t, http.StatusOK, girlsResponse)

	for range 2 {
		shows, err := SearchShows(context.Background(), "girls")
		if err != nil {
			t.Fatal(err)
		}

		show := shows[0]
		if show.Summary != "Hannah & friends" || *show.Year != "2012" || *show.Rating != "7.0" ||
			*show.Network != "HBO" || show.Image != nil {
			t.Fatalf("неверное приведение ответа TVmaze: %+v", show)
		}
	}

	if *requests != 1 {
		t.Fatalf("повторный запрос в течение 5 минут берётся из кеша, запросов к TVmaze: %d", *requests)
	}
}

func TestErrorsAreNotCached(t *testing.T) {
	requests := useFakeTVmaze(t, http.StatusTooManyRequests, "")

	for range 2 {
		if _, err := GetShow(context.Background(), 1); !errors.Is(err, ErrRateLimit) {
			t.Fatalf("429 от TVmaze — ErrRateLimit, получили %v", err)
		}
	}

	if *requests != 2 {
		t.Fatalf("ошибки не кешируются, запросов к TVmaze: %d", *requests)
	}
}

func useFakeTVmaze(t *testing.T, status int, body string) *int {
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requests++
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))

	previousURL, previousCache := baseURL, cache
	baseURL, cache = server.URL, newCacheStore(time.Now)

	t.Cleanup(func() {
		server.Close()
		baseURL, cache = previousURL, previousCache
	})

	return &requests
}
