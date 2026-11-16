package shows

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/wildbober228/it-sinyak-homework_1_backend/web"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/shows", handleSearch)
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	page, err := parsePage(params.Get("page"))
	if err != nil {
		web.WriteError(w, err)

		return
	}

	query := searchQuery{Text: params.Get("q"), Genre: params.Get("genre"), Page: page}

	result, err := searchShows(r.Context(), query)
	if err != nil {
		web.WriteError(w, err)

		return
	}

	web.WriteJSON(w, http.StatusOK, result)
}

func parsePage(value string) (int, error) {
	if value == "" {
		return 1, nil
	}

	page, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%w: page %q", web.ErrBadRequest, value)
	}

	return page, nil
}
