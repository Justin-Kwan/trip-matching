package redis

import (
	"log"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"order-matching/internal/config"
)

const (
	_floatMaxDelta = 0.01001
)

var (
	_geoDB *GeoDB

	// test points of interest
	testPOIs = []TestPOI{
		TestPOI{
			id: "test_id1",
			coord: map[string]float64{
				"lon": 1,
				"lat": 2,
			},
		},
		TestPOI{
			id: "test_id2",
			coord: map[string]float64{
				"lon": 3.141878,
				"lat": 5.123,
			},
		},
		// edge case (max lon and max lat)
		TestPOI{
			id: "test_id3",
			coord: map[string]float64{
				"lon": 180,
				"lat": 85.05112878,
			},
		},
		// edge case (min lon and min lat)
		TestPOI{
			id: "test_id4",
			coord: map[string]float64{
				"lon": -180,
				"lat": -85.05112878,
			},
		},
		// edge case (max lon and min lat)
		TestPOI{
			id: "test_id5",
			coord: map[string]float64{
				"lon": 180,
				"lat": -85.05112878,
			},
		},
		// edge case (min lon and max lat)
		TestPOI{
			id: "test_id6",
			coord: map[string]float64{
				"lon": -180,
				"lat": 85.05112878,
			},
		},
		TestPOI{
			id: "test_id7",
			coord: map[string]float64{
				"lon": 0,
				"lat": 0,
			},
		},
		TestPOI{
			id: "test_id8",
			coord: map[string]float64{
				"lon": 45.12321,
				"lat": 32.124,
			},
		},
		TestPOI{
			id: "test_id9",
			coord: map[string]float64{
				"lon": 75.987567,
				"lat": 67.124122124,
			},
		},
		TestPOI{
			id: "test_id10",
			coord: map[string]float64{
				"lon": 75.987,
				"lat": 67.12412,
			},
		},
		TestPOI{
			id: "test_id11",
			coord: map[string]float64{
				"lon": 45.213,
				"lat": 74.98723,
			},
		},
		TestPOI{
			id: "test_id12",
			coord: map[string]float64{
				"lon": 45.213,
				"lat": 74.98723,
			},
		},
	}
)

type TestPOI struct {
	id    string
	coord map[string]float64
}

type geoDBTestConstants struct {
	configFilePath string
	env            string
	setIndex       string
}

func newGeoDBTestConstants() *geoDBTestConstants {
	return &geoDBTestConstants{
		configFilePath: "../../../",
		env:            "test",
		setIndex:       "test_index",
	}
}

func setupGeoDBTests() func() {
	tc := newGeoDBTestConstants()

	cfg, _ := config.NewConfig(tc.configFilePath, tc.env)
	geoDBPool := NewPool(&(*cfg).RedisGeoDB)

	_geoDB = NewGeoDB(geoDBPool, tc.setIndex)
	_geoDB.Clear()

	return func() {
		_geoDB.Clear()
	}
}

func TestInsert(t *testing.T) {
	teardownTests := setupGeoDBTests()
	defer teardownTests()

	for _, testPOI := range testPOIs {
		// function under test
		if err := _geoDB.Insert(testPOI.id, testPOI.coord); err != nil {
			log.Fatalf(err.Error())
		}

		// select and assert point of interest exists
		coord, err := _geoDB.Select(testPOI.id)
		if err != nil {
			teardownTests()
			log.Fatalf(err.Error())
		}

		assert.InDelta(t, testPOI.coord["lon"], coord["lon"], _floatMaxDelta)
		assert.InDelta(t, testPOI.coord["lat"], coord["lat"], _floatMaxDelta)
	}
}

func TestSelect(t *testing.T) {
	teardownTests := setupGeoDBTests()
	defer teardownTests()

	var errTestCases = []struct {
		inputKeyId  string
		expectedErr error
	}{
		{"non_existent_key", errors.Errorf("Error selecting POI with key 'non_existent_key'")},
		{" ", errors.Errorf("Error selecting POI with key ' '")},
		{"", errors.Errorf("Error selecting POI with key ''")},
		{"*", errors.Errorf("Error selecting POI with key '*'")},
	}

	for _, testCase := range errTestCases {
		_, err := _geoDB.Select(testCase.inputKeyId)
		assert.EqualError(t, err, testCase.expectedErr.Error())
	}
}

func TestSelectNearestInRadius(t *testing.T) {
	teardownTests := setupGeoDBTests()
	defer teardownTests()

	// test on empty db
	coord := map[string]float64{53.123, 84.9823}
	POIkeyId, err := _geoDB.SelectNearestInRadius(coord, 10000)

	
	assert.Equal(t, "", POIkeyId, "should match closest point of interest's key id within radius")

	// setup
	for _, testPOI := range testPOIs {
		if err := _geoDB.Insert(testPOI.id, testPOI.coord); err != nil {
			log.Fatalf(err.Error())
		}
	}

	var testCases = []struct {
		inputCoord    map[string]float64
		inputRadius   float64
		expectedKeyId string
		expectedErr   error
	}{
		{
			map[string]float64{"lon": 90, "lat": 65},
			1,
			"",
			errors.Errorf("Error selecting nearest POI to (90, 65) within 1 km"),
		},
		{
			map[string]float64{"lon": 86.3234, "lat": 66.123},
			0,
			"",
			errors.Errorf("Error selecting nearest POI to (86.3234, 66.123) within 0 km"),
		},
		{
			map[string]float64{"lon": 90, "lat": 65},
			-1,
			"",
			errors.Errorf("Error selecting nearest POI to (90, 65) within -1 km"),
		},
		{
			map[string]float64{"lon": 75.8, "lat": 67.124},
			8.086,
			"test_id10",
			nil,
		},
		{
			map[string]float64{"lon": 75.8, "lat": 67.124},
			20,
			"test_id10",
			nil,
		},
		{
			map[string]float64{"lon": 4, "lat": 5},
			100,
			"test_id2",
			nil,
		},
		{
			map[string]float64{"lon": 45, "lat": 32},
			19,
			"test_id8",
			nil,
		},
		{
			map[string]float64{"lon": 46.2, "lat": 75.001},
			29,
			"test_id11",
			nil,
		},
		{
			map[string]float64{"lon": -178.8991238, "lat": -80.2312431},
			536.303,
			"test_id4",
			nil,
		},
	}

	for _, testCase := range testCases {
		// function under test
		POIkeyId, err := _geoDB.SelectNearestInRadius(testCase.inputCoord, testCase.inputRadius)
		assert.Equal(t, testCase.expectedKeyId, POIkeyId, "should match closest point of interest's key id within radius")

		if testCase.expectedErr != nil {
			assert.EqualError(t, err, testCase.expectedErr.Error(), "should assert error is returned")
		} else {
			assert.NoError(t, err, "should assert no error is returned")
		}
	}
}

func TestDelete(t *testing.T) {
	teardownTests := setupGeoDBTests()
	defer teardownTests()

	for _, testPOI := range testPOIs {
		// setup
		if err := _geoDB.Insert(testPOI.id, testPOI.coord); err != nil {
			log.Fatalf(err.Error())
		}

		// select and assert point of interest exists before deleting
		coord, err := _geoDB.Select(testPOI.id)
		if err != nil {
			teardownTests()
			log.Fatalf(err.Error())
		}

		assert.InDelta(t, testPOI.coord["lon"], coord["lon"], _floatMaxDelta)
		assert.InDelta(t, testPOI.coord["lat"], coord["lat"], _floatMaxDelta)

		// function under test
		if err := _geoDB.Delete(testPOI.id); err != nil {
			log.Fatalf(err.Error())
		}

		// assert point of interest is deleted
		_, err = _geoDB.Select(testPOI.id)
		assert.EqualError(t, err, "Error selecting POI with key '"+testPOI.id+"'")
	}

	var errTestCases = []struct {
		inputKeyId  string
		expectedErr error
	}{
		{"non_existent_key", errors.Errorf("Error deleting POI with key 'non_existent_key'")},
		{" ", errors.Errorf("Error deleting POI with key ' '")},
		{"", errors.Errorf("Error deleting POI with key ''")},
		{"*&^", errors.Errorf("Error deleting POI with key '*&^'")},
	}

	for _, testCase := range errTestCases {
		err := _geoDB.Delete(testCase.inputKeyId)
		assert.EqualError(t, err, testCase.expectedErr.Error())
	}
}
