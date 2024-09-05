package mtg

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test for FetchCardsFromAPI
func TestFetchCardsFromAPI(t *testing.T) {
    err := FetchCardsFromAPI()
    assert.NoError(t, err, "Expected no error while fetching cards")
}

// Mock HTTP response for better test coverage
func TestFetchCardsFromAPIMock(t *testing.T) {
    mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintln(w, `{"cards": [{"id": "1", "name": "Test Card"}]}`)
    }))
    defer mockServer.Close()

    err := FetchCardsFromAPI(mockServer.URL)
    assert.NoError(t, err, "Expected no error with mock server")
}
