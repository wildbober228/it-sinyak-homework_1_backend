package tvmaze

import "errors"

var (
	ErrNotFound    = errors.New("tvmaze: not found")
	ErrRateLimit   = errors.New("tvmaze: rate limit")
	ErrUnavailable = errors.New("tvmaze: unavailable")
)
