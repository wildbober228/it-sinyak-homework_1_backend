package health

import (
	"net/http"

	"github.com/wildbober228/it-sinyak-homework_1_backend/web"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/health", handleHealth)
}

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	web.WriteJSON(w, http.StatusOK, healthStatus{Status: "ok"})
}
