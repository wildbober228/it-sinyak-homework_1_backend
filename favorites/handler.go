package favorites

import (
	"net/http"

	"github.com/wildbober228/it-sinyak-homework_1_backend/tvmaze"
	"github.com/wildbober228/it-sinyak-homework_1_backend/web"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/favorites", handleList)
	mux.HandleFunc("PUT /api/favorites/{id}", handleAdd)
	mux.HandleFunc("DELETE /api/favorites", handleRemove)
}

func handleList(w http.ResponseWriter, r *http.Request) {
	clientID, err := readClientID(r)
	if err != nil {
		web.WriteError(w, err)

		return
	}

	writeList(w, listFavorites(clientID))
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	clientID, err := readClientID(r)
	if err != nil {
		web.WriteError(w, err)

		return
	}

	showID, err := parseID(r.PathValue("id"))
	if err != nil {
		web.WriteError(w, err)

		return
	}

	shows, err := addFavorite(r.Context(), clientID, showID)
	if err != nil {
		web.WriteError(w, err)

		return
	}

	writeList(w, shows)
}

func handleRemove(w http.ResponseWriter, r *http.Request) {
	clientID, err := readClientID(r)
	if err != nil {
		web.WriteError(w, err)

		return
	}

	showIDs, err := parseIDs(r.URL.Query().Get("ids"))
	if err != nil {
		web.WriteError(w, err)

		return
	}

	writeList(w, removeFavorites(clientID, showIDs))
}

func writeList(w http.ResponseWriter, shows []tvmaze.Show) {
	web.WriteJSON(w, http.StatusOK, favoritesList{Items: shows})
}
