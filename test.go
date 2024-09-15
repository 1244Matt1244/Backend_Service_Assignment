package mtg

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDB is a mock database structure for testing
type MockDB struct {
	mock.Mock
}

// NamedExec is a mock function that simulates database execution
func (m *MockDB) NamedExec(query string, arg interface{}) (sqlx.Result, error) {
	args := m.Called(query, arg)
	return nil, args.Error(1)
}

// TestFetchCardsFromAPIMock tests the FetchCardsFromAPI function with a mock server
func TestFetchCardsFromAPIMock(t *testing.T) {
	// Create a mock server that returns a fixed response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"cards": [{"id": "1", "name": "Test Card", "colors": ["red"], "cmc": 3}]}`)
	}))
	defer mockServer.Close()

	// Mock the database
	db := new(MockDB)
	db.On("NamedExec", mock.Anything, mock.Anything).Return(nil, nil)

	// Call the function under test
	err := FetchCardsFromAPI(mockServer.URL, db)

	// Assert that there was no error
	assert.NoError(t, err, "Expected no error with mock server")

	// Additional assertions to verify that the function behaves correctly
	// Check if the NamedExec method was called with the expected arguments
	db.AssertCalled(t, "NamedExec", mock.Anything, mock.Anything)
}
