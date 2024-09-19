package camera

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestFindCamerasWithinRadius(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer dbMock.Close()

	// Flexible matching for the SQL query
	query := regexp.QuoteMeta(`SELECT id, name, latitude, longitude FROM traffic_cameras 
        WHERE ST_DWithin(ST_SetSRID(ST_MakePoint($1, $2), 4326), 
        ST_SetSRID(ST_MakePoint(longitude, latitude), 4326), $3)`)

	rows := sqlmock.NewRows([]string{"id", "name", "latitude", "longitude"}).
		AddRow("1", "Camera 1", 45.0, 15.0).
		AddRow("2", "Camera 2", 45.1, 15.1)

	mock.ExpectQuery(query).WithArgs(15.0, 45.0, 1000.0).WillReturnRows(rows)

	cameras, err := FindCamerasWithinRadius(dbMock, 45.0, 15.0, 1000.0)
	assert.NoError(t, err)
	assert.NotNil(t, cameras)
	assert.Len(t, cameras, 2)
	assert.Equal(t, "1", cameras[0].ID)
	assert.Equal(t, "Camera 1", cameras[0].Name)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestFindCamerasWithinRadius_NoResults(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer dbMock.Close()

	query := regexp.QuoteMeta(`SELECT id, name, latitude, longitude FROM traffic_cameras 
        WHERE ST_DWithin(ST_SetSRID(ST_MakePoint($1, $2), 4326), 
        ST_SetSRID(ST_MakePoint(longitude, latitude), 4326), $3)`)

	mock.ExpectQuery(query).WithArgs(15.0, 45.0, 1000.0).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "latitude", "longitude"}))

	cameras, err := FindCamerasWithinRadius(dbMock, 45.0, 15.0, 1000.0)
	assert.NoError(t, err)
	assert.NotNil(t, cameras)
	assert.Len(t, cameras, 0)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestFindCamerasWithinRadius_InvalidArgs(t *testing.T) {
	dbMock, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer dbMock.Close()

	_, err = FindCamerasWithinRadius(dbMock, 45.0, 15.0, -1000.0)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid radius: must be greater than zero")
}

func TestGetCameraByID_InvalidID(t *testing.T) {
	dbMock, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer dbMock.Close()

	_, err = GetCameraByID(dbMock, "")
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid camera ID")
}
