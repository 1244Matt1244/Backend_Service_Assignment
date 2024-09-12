package camera

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCameraByID(t *testing.T) {
	// Test with existing camera
	cameras := []Camera{
		{ID: "1", Latitude: 40.7128, Longitude: -74.0060},
	}

	camera, err := GetCameraByID("1", cameras)
	assert.NoError(t, err)
	assert.Equal(t, 40.7128, camera.Latitude)
	assert.Equal(t, -74.0060, camera.Longitude)

	// Test with non-existent camera
	camera, err = GetCameraByID("999", cameras)
	assert.Error(t, err)
	assert.Equal(t, "camera with ID 999 not found", err.Error())
	assert.Equal(t, Camera{}, camera)

	// Test with empty camera list
	cameras = []Camera{}
	camera, err = GetCameraByID("1", cameras)
	assert.Error(t, err)
	assert.Equal(t, "camera with ID 1 not found", err.Error())
	assert.Equal(t, Camera{}, camera)
}
