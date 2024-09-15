package camera

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDB is a mock database structure for testing
type MockDB struct {
	mock.Mock
}

// InsertCameras simulates inserting cameras into the database
func (m *MockDB) InsertCameras(cameras []Camera) error {
	args := m.Called(cameras)
	return args.Error(0)
}

// TestInsertCameras tests the InsertCameras function
func TestInsertCameras(t *testing.T) {
	// Mock database connection setup
	db := new(MockDB)
	cameras := []Camera{{ID: "1", Latitude: 40.7128, Longitude: -74.0060}}

	// Mock the InsertCameras call
	db.On("InsertCameras", cameras).Return(nil)

	// Simulate InsertCameras
	err := InsertCameras(db, cameras)
	assert.NoError(t, err)

	// Test for GetCameraByID function with existing camera
	camera, err := GetCameraByID("1", cameras)
	assert.NoError(t, err)
	assert.Equal(t, 40.7128, camera.Latitude)
	assert.Equal(t, -74.0060, camera.Longitude)

	// Test with non-existent camera
	camera, err = GetCameraByID("999", cameras)
	assert.Error(t, err)
	assert.Equal(t, "camera with ID 999 not found", err.Error())
}

// GetCameraByID is a helper function to fetch a camera by ID
func GetCameraByID(id string, cameras []Camera) (Camera, error) {
	for _, camera := range cameras {
		if camera.ID == id {
			return camera, nil
		}
	}
	return Camera{}, errors.New("camera with ID " + id + " not found")
}

// InsertCameras inserts the provided cameras into the mock database
func InsertCameras(db *MockDB, cameras []Camera) error {
	return db.InsertCameras(cameras)
}
