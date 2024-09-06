// camera_service.go

package camera

import "fmt"

// GetCameraByID retrieves a camera by its ID
func GetCameraByID(id string, cameras []Camera) (*Camera, error) {
	for _, camera := range cameras {
		if camera.ID == id {
			return &camera, nil
		}
	}
	return nil, fmt.Errorf("camera not found")
}
