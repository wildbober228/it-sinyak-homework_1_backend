package favorites

import (
	"testing"

	"github.com/wildbober228/it-sinyak-homework_1_backend/tvmaze"
)

func ids(shows []tvmaze.Show) []int {
	result := make([]int, 0, len(shows))

	for _, show := range shows {
		result = append(result, show.ID)
	}

	return result
}

func TestStoreListsNewestFirstWithoutDuplicates(t *testing.T) {
	favorites := newFavoritesStore()

	favorites.add("alice", tvmaze.Show{ID: 1})
	favorites.add("alice", tvmaze.Show{ID: 2})
	favorites.add("alice", tvmaze.Show{ID: 1})

	if got := ids(favorites.list("alice")); len(got) != 2 || got[0] != 2 || got[1] != 1 {
		t.Fatalf("свежие сверху, без повторов: %v", got)
	}
}

func TestStoreKeepsClientsApart(t *testing.T) {
	favorites := newFavoritesStore()

	favorites.add("alice", tvmaze.Show{ID: 1})
	favorites.add("bob", tvmaze.Show{ID: 2})
	favorites.remove("bob", []int{2, 3})

	if got := ids(favorites.list("alice")); len(got) != 1 || got[0] != 1 {
		t.Fatalf("у каждого браузера своё избранное: %v", got)
	}

	if got := favorites.list("bob"); len(got) != 0 || got == nil {
		t.Fatalf("после удаления — пустой список, а не null: %v", got)
	}

	if _, ok := favorites.clients["bob"]; ok {
		t.Fatal("клиент без избранного не занимает память")
	}
}
