package web

import (
	"net/http"
	"slices"
)

func CORS(origins []string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		origin := r.Header.Get("Origin")

		header.Add("Vary", "Origin")

		if slices.Contains(origins, origin) {
			header.Set("Access-Control-Allow-Origin", origin)
		}

		if r.Method == http.MethodOptions {
			header.Set("Access-Control-Allow-Methods", allowedMethods)
			header.Set("Access-Control-Allow-Headers", allowedHeaders)
			header.Set("Access-Control-Max-Age", preflightMaxAge)
			w.WriteHeader(http.StatusNoContent)

			return
		}

		next.ServeHTTP(w, r)
	})
}
