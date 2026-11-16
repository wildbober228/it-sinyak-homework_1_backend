package web

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/wildbober228/it-sinyak-homework_1_backend/tvmaze"
)

func WriteJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("write response: %v", err)
	}
}

func WriteError(w http.ResponseWriter, err error) {
	status, code := classify(err)

	if status >= http.StatusInternalServerError {
		log.Printf("request failed: %v", err)
	}

	WriteJSON(w, status, errorBody{Error: errorDetails{Code: code}})
}

func classify(err error) (int, string) {
	switch {
	case errors.Is(err, ErrBadRequest):
		return http.StatusBadRequest, "badRequest"
	case errors.Is(err, tvmaze.ErrNotFound):
		return http.StatusNotFound, "notFound"
	case errors.Is(err, tvmaze.ErrRateLimit):
		return http.StatusTooManyRequests, "rateLimit"
	case errors.Is(err, tvmaze.ErrUnavailable):
		return http.StatusBadGateway, "upstream"
	default:
		return http.StatusInternalServerError, "internal"
	}
}
