package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCORSAnswersPreflightForAllowedOrigin(t *testing.T) {
	handler := CORS([]string{"https://example.github.io"}, http.NotFoundHandler())
	request := httptest.NewRequest(http.MethodOptions, "/api/favorites/1", nil)
	request.Header.Set("Origin", "https://example.github.io")
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, request)

	header := recorder.Header()
	if recorder.Code != http.StatusNoContent ||
		header.Get("Access-Control-Allow-Origin") != "https://example.github.io" ||
		header.Get("Access-Control-Allow-Headers") != allowedHeaders {
		t.Fatalf("preflight не разрешён: %d %v", recorder.Code, header)
	}
}

func TestCORSIgnoresUnknownOrigin(t *testing.T) {
	handler := CORS([]string{"https://example.github.io"}, http.NotFoundHandler())
	request := httptest.NewRequest(http.MethodGet, "/api/shows", nil)
	request.Header.Set("Origin", "https://evil.example")
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, request)

	if origin := recorder.Header().Get("Access-Control-Allow-Origin"); origin != "" {
		t.Fatalf("чужой origin не разрешаем, получили %q", origin)
	}
}
