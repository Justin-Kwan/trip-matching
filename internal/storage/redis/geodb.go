package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

type GeoDB struct {
	pool  *redis.Pool
	dbNum int
	index string
}

func NewGeoDB(pool *redis.Pool, index string) *GeoDB {
	return &GeoDB{
		pool:  pool,
		index: index,
	}
}

func (db *GeoDB) Insert(keyId string, coord map[string]float64) error {
	conn := db.pool.Get()
	defer conn.Close()

	_, err := conn.Do("GEOADD", db.index, coord["lon"], coord["lat"], keyId)

	if err != nil {
		return errors.Errorf("Error adding POI with key '%s': %v", keyId, err)
	}

	return nil
}

func (db *GeoDB) Select(keyId string) (map[string]float64, error) {
	conn := db.pool.Get()
	defer conn.Close()

	res, err := redis.Positions(conn.Do("GEOPOS", db.index, keyId))

	keyNotFound := res[0] == nil
	if err != nil || keyNotFound {
		return nil, errors.Errorf("Error selecting POI with key '%s'", keyId)
	}

	coord := map[string]float64{
		"lon": res[0][0],
		"lat": res[0][1],
	}

	return coord, nil
}

func (db *GeoDB) SelectNearestInRadius(coords map[string]float64, radius float64) (string, error) {
	conn := db.pool.Get()
	defer conn.Close()

	res, err := redis.Strings(conn.Do(
		"GEORADIUS",
		db.index,
		coords["lon"],
		coords["lat"],
		radius,
		"km",
		"ASC",
	))

	noNearbyPOI := len(res) == 0
	if err != nil || noNearbyPOI {
		errStr := "Error selecting nearest POI to (%v, %v) within %v km"
		return "", errors.Errorf(errStr, coords["lon"], coords["lat"], radius)
	}

	closestPOIKeyId := res[0]
	return closestPOIKeyId, nil
}

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

func (db *GeoDB) Clear() error {
	conn := db.pool.Get()
	defer conn.Close()

	if _, err := conn.Do("FLUSHDB"); err != nil {
		return errors.Errorf("Error clearing all key resue pairs: %v", err)
	}

	return nil
}
