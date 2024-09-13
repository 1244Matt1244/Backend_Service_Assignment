package camera

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertCameras(t *testing.T) {
	// Mock database connection setup
	db := new(MockDB)
	cameras := []Camera{{ID: "1", Latitude: 40.7128, Longitude: -74.0060}}

	// Simulate InsertCameras
	db.On("InsertCameras", cameras).Return(nil)

	err := InsertCameras(db, cameras)
	assert.NoError(t, err)

	// Test for GetCameraByID function
	camera, err := GetCameraByID("1", cameras)
	assert.NoError(t, err)
	assert.Equal(t, 40.7128, camera.Latitude)
	assert.Equal(t, -74.0060, camera.Longitude)

	// Test with non-existent camera
	camera, err = GetCameraByID("999", cameras)
	assert.Error(t, err)
	assert.Equal(t, "camera with ID 999 not found", err.Error())
}
