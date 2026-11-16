package favorites

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/wildbober228/it-sinyak-homework_1_backend/web"
)

var clientIDPattern = regexp.MustCompile(clientIDFormat)

func readClientID(r *http.Request) (string, error) {
	clientID := r.Header.Get(clientIDHeader)

	if !clientIDPattern.MatchString(clientID) {
		return "", fmt.Errorf("%w: client id %q", web.ErrBadRequest, clientID)
	}

	return clientID, nil
}

func parseIDs(value string) ([]int, error) {
	parts := strings.Split(value, ",")

	if len(parts) > maxRemoveIDs {
		return nil, fmt.Errorf("%w: too many ids", web.ErrBadRequest)
	}

	ids := make([]int, 0, len(parts))

	for _, part := range parts {
		id, err := parseID(part)
		if err != nil {
			return nil, err
		}

		ids = append(ids, id)
	}

	return ids, nil
}

func parseID(value string) (int, error) {
	id, err := strconv.Atoi(strings.TrimSpace(value))

	if err != nil || id <= 0 {
		return 0, fmt.Errorf("%w: show id %q", web.ErrBadRequest, value)
	}

	return id, nil
}
