package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/wildbober228/it-sinyak-homework_1_backend/favorites"
	"github.com/wildbober228/it-sinyak-homework_1_backend/health"
	"github.com/wildbober228/it-sinyak-homework_1_backend/shows"
	"github.com/wildbober228/it-sinyak-homework_1_backend/web"
)

func main() {
	mux := http.NewServeMux()
	health.Register(mux)
	shows.Register(mux)
	favorites.Register(mux)

	origins := strings.Split(envOr("ALLOWED_ORIGINS", defaultOrigins), ",")
	server := &http.Server{
		Addr:              ":" + envOr("PORT", defaultPort),
		Handler:           web.CORS(origins, mux),
		ReadHeaderTimeout: readHeaderTimeout,
	}

	log.Printf("listening on %s, allowed origins %v", server.Addr, origins)
	log.Fatal(server.ListenAndServe())
}

func envOr(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
