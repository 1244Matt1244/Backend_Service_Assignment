package mtg

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestFetchCardsFromAPIMock tests the FetchCardsFromAPI function with a mock server
func TestFetchCardsFromAPIMock(t *testing.T) {
	// Create a mock server that returns a fixed response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"cards": [{"id": "1", "name": "Test Card"}]}`)
	}))
	defer mockServer.Close()

	// Call the function under test
	err := FetchCardsFromAPI(mockServer.URL)

	// Assert that there was no error
	assert.NoError(t, err, "Expected no error with mock server")

	// Additional assertions to verify that the function behaves correctly
	// For example, check if the fetched data matches expected results
	// You might need to modify FetchCardsFromAPI to return data for further assertions
}
