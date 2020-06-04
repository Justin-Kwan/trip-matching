package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

type GeoDB struct {
	pool   *redis.Pool
	locker *Locker
	index  string
}

// Returns a new geo database access struct given a redis connection
// pool and an index name to store all longitude, latitude pairs
// within the key values store.
func NewGeoDB(pool *redis.Pool, index string) *GeoDB {
	return &GeoDB{
		pool:   pool,
		index:  index,
	}
}

// Adds a point of interest into the geo spatial database, given
// a key id and a map containing the longtitude and latitude
// coordinates. A generic error is returned for all errors.
func (db *GeoDB) Insert(keyId string, coord map[string]float64) error {
	conn := db.pool.Get()
	defer conn.Close()

	_, err := conn.Do("GEOADD", db.index, coord["lon"], coord["lat"], keyId)

	if err != nil {
		return errors.Errorf("Error adding POI with key '%s': %v", keyId, err)
	}

	return nil
}

// Returns a map containing the longitude and latitude coordinates
// of a point of interest given its key id. A generic error is
// returned for all errors.
func (db *GeoDB) Select(keyId string) (map[string]float64, error) {
	conn := db.pool.Get()
	defer conn.Close()

	res, err := redis.Positions(conn.Do("GEOPOS", db.index, keyId))

	emptyCoord := res[0] == nil
	if err != nil || emptyCoord {
		return nil, errors.Errorf("Error selecting POI with key '%s'", keyId)
	}

	coord := map[string]float64{
		"lon": res[0][0],
		"lat": res[0][1],
	}

	return coord, nil
}

func (db *GeoDB) SelectAllInRadius(coords map[string]float64, radius float64) ([]string, error) {
	conn := db.pool.Get()
	defer conn.Close()

	POIKeyIds, err := redis.Strings(conn.Do(
		"GEORADIUS",
		db.index,
		coords["lon"],
		coords["lat"],
		radius,
		"km",
		"ASC",
	))

	if err != nil {
		errStr := "Error selecting nearest POI to (%v, %v) within %v km"
		return nil, errors.Errorf(errStr, coords["lon"], coords["lat"], radius)
	}

	return POIKeyIds, nil
}

// Deletes a specific point of interest (lon, lat) by key in the
// current geospacial index. A generic error is returned for all other
// errors.
func (db *GeoDB) Delete(keyId string) error {
	conn := db.pool.Get()
	defer conn.Close()

	res, err := redis.Bool(conn.Do("ZREM", db.index, keyId))

	keyNotFound := res == false
	if err != nil || keyNotFound {
		return errors.Errorf("Error deleting POI with key '%s'", keyId)
	}

	return nil
}

// Clears the entire key value store. A generic error is returned for
// all errors.
func (db *GeoDB) Clear() error {
	conn := db.pool.Get()
	defer conn.Close()

	if _, err := conn.Do("FLUSHDB"); err != nil {
		return errors.Errorf("Error clearing all key resue pairs: %v", err)
	}

	return nil
}
