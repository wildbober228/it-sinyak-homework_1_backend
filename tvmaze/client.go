package tvmaze

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

var (
	baseURL    = "https://api.tvmaze.com"
	httpClient = &http.Client{Timeout: requestTimeout}
)

func SearchShows(ctx context.Context, query string) ([]Show, error) {
	var results []apiSearchResult

	if err := get(ctx, "/search/shows?q="+url.QueryEscape(query), &results); err != nil {
		return nil, err
	}

	shows := make([]Show, 0, len(results))

	for _, result := range results {
		shows = append(shows, toShow(result.Show))
	}

	return shows, nil
}

func GetShow(ctx context.Context, id int) (Show, error) {
	var show apiShow

	if err := get(ctx, "/shows/"+strconv.Itoa(id), &show); err != nil {
		return Show{}, err
	}

	return toShow(show), nil
}

func get(ctx context.Context, path string, target any) error {
	address := baseURL + path

	if body, ok := cache.get(address); ok {
		log.Printf("tvmaze cache hit %s", path)

		return json.Unmarshal(body, target)
	}

	log.Printf("tvmaze fetch %s", path)

	body, err := fetch(ctx, address)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("%w: %w", ErrUnavailable, err)
	}

	cache.set(address, body)

	return nil
}

func fetch(ctx context.Context, address string) ([]byte, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, address, nil)
	if err != nil {
		return nil, err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrUnavailable, err)
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		return io.ReadAll(response.Body)
	case http.StatusNotFound:
		return nil, ErrNotFound
	case http.StatusTooManyRequests:
		return nil, ErrRateLimit
	default:
		return nil, fmt.Errorf("%w: status %d", ErrUnavailable, response.StatusCode)
	}
}
