// camera_service.go

package camera

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Camera struct {
	ID        string  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func ParseCameras(filename string) ([]Camera, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cameras []Camera
	reader := csv.NewReader(file)
	_, err = reader.Read() // Skip header
	if err != nil {
		return nil, err
	}

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		latitude, _ := strconv.ParseFloat(record[1], 64)
		longitude, _ := strconv.ParseFloat(record[2], 64)
		cameras = append(cameras, Camera{
			ID:        record[0],
			Latitude:  latitude,
			Longitude: longitude,
		})
	}

	return cameras, nil
}

func GetCameraByID(id string, cameras []Camera) (Camera, error) {
	for _, camera := range cameras {
		if camera.ID == id {
			return camera, nil
		}
	}
	return Camera{}, fmt.Errorf("camera not found")
}
