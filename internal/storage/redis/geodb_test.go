package redis

import (
	"log"
	"testing"

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
	}
)

type TestPOI struct {
	id    string
	coord map[string]float64
}

type geoDBTestConstants struct {
	configFilePath string
	env            string
	dbNum          int
	setIndex       string
}

func newGeoDBTestConstants() *geoDBTestConstants {
	return &geoDBTestConstants{
		configFilePath: "../../../",
		env:            "test",
		dbNum:          1,
		setIndex:       "test_index",
	}
}

func setupGeoDBTests() func() {
	tc := newGeoDBTestConstants()

	testCfg, _ := config.NewConfig(tc.configFilePath, tc.env)
	testRedisCfg := &(*testCfg).Redis
	redisPool, _ := NewPool(testRedisCfg)

	_geoDB = NewGeoDB(redisPool, tc.dbNum, tc.setIndex)

	return func() {
		_geoDB.Clear()
	}
}

func TestInsert(t *testing.T) {
	teardownGeoDBTests := setupGeoDBTests()
	defer teardownGeoDBTests()

	for _, testPOI := range testPOIs {
		// function under test
		if err := _geoDB.Insert(testPOI.id, testPOI.coord); err != nil {
			log.Fatalf(err.Error())
		}

		// select and assert point of interest exists
		coord, err := _geoDB.Select(testPOI.id)
		if err != nil {
			log.Fatalf(err.Error())
		}

		assert.InDelta(t, testPOI.coord["lon"], coord["lon"], _floatMaxDelta)
		assert.InDelta(t, testPOI.coord["lat"], coord["lat"], _floatMaxDelta)
	}
}

func TestSelect(t *testing.T) {
	teardownGeoDBTests := setupGeoDBTests()
	defer teardownGeoDBTests()

	// assert errors when selecting non-existent keys
	_, err := _geoDB.Select("non_existent_key")
	assert.EqualError(t, err, "Error selecting POI with key 'non_existent_key'")

	_, err = _geoDB.Select(" ")
	assert.EqualError(t, err, "Error selecting POI with key ' '")

	_, err = _geoDB.Select("")
	assert.EqualError(t, err, "Error selecting POI with key ''")

	_, err = _geoDB.Select("%")
	assert.EqualError(t, err, "Error selecting POI with key '%'")
}

func TestDelete(t *testing.T) {
	teardownGeoDBTests := setupGeoDBTests()
	defer teardownGeoDBTests()

	for _, testPOI := range testPOIs {
		// setup
		if err := _geoDB.Insert(testPOI.id, testPOI.coord); err != nil {
			log.Fatalf(err.Error())
		}

		// select and assert point of interest exists before deleting
		coord, err := _geoDB.Select(testPOI.id)
		if err != nil {
			log.Fatalf(err.Error())
		}
		
		assert.InDelta(t, testPOI.coord["lon"], coord["lon"], _floatMaxDelta)
		assert.InDelta(t, testPOI.coord["lat"], coord["lat"], _floatMaxDelta)

		// function under test
		if err := _geoDB.Delete(testPOI.id); err != nil {
			log.Fatalf(err.Error())
		}

		// assert point of interest is deleted
		coord, err = _geoDB.Select(testPOI.id)
		assert.EqualError(t, err, "Error selecting POI with key '"+testPOI.id+"'")
	}
}
