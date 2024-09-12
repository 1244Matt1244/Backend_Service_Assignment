package camera

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertCameras(t *testing.T) {
    db := // setup mock db
    cameras := []Camera{{ID: "1", Name: "Cam1", Latitude: 1.234, Longitude: 5.678}}
    err := InsertCameras(db, cameras)
    if err != nil {
        t.Errorf("Expected no error, but got %v", err)
    }
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

