package camera

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCameraByID(t *testing.T) {
	cameras := []Camera{
		{ID: "1", Latitude: 40.7128, Longitude: -74.0060},
	}

	camera, err := GetCameraByID("1", cameras)
	assert.NoError(t, err)
	assert.Equal(t, 40.7128, camera.Latitude)
	assert.Equal(t, -74.0060, camera.Longitude)
}
