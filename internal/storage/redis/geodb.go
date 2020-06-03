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

func NewGeoDB(pool *redis.Pool, dbNum int, index string) *GeoDB {
	return &GeoDB{
		pool:  pool,
		dbNum: dbNum,
		index: index,
	}
}

func (db *GeoDB) Insert(keyId string, coord map[string]float64) error {
	conn := db.pool.Get()
	conn.Do("SELECT", db.dbNum)
	defer conn.Close()

	_, err := conn.Do(
		"GEOADD",
		db.index,
		coord["lon"],
		coord["lat"],
		keyId,
	)

	if err != nil {
		return errors.Errorf("Error adding POI with key '%s': %v", keyId, err)
	}

	return nil
}

func (db *GeoDB) Select(keyId string) (map[string]float64, error) {
	conn := db.pool.Get()
	conn.Do("SELECT", db.dbNum)
	defer conn.Close()

	val, err := redis.Positions(conn.Do(
		"GEOPOS",
		db.index,
		keyId,
	))

	if err != nil || val[0] == nil {
		return nil, errors.Errorf("Error selecting POI with key '%s'", keyId)
	}

	coord := map[string]float64{
		"lon": val[0][0],
		"lat": val[0][1],
	}

	return coord, nil
}

// will have to lock the key returned
func (db *GeoDB) SelectNearestInRadius(coords map[string]float64, radius float64) (string, error) {
	conn := db.pool.Get()
	conn.Do("SELECT", db.dbNum)
	defer conn.Close()

	val, err := redis.Strings(conn.Do(
		"GEORADIUS",
		db.index,
		coords["lon"],
		coords["lat"],
		radius,
		"km",
		"ASC",
	))

	if err != nil || len(val) == 0 {
		return "", errors.Errorf("Error selecting nearest POI within %v km", radius)
	}

	closestPOIKeyId := val[0]

	// aquire lock on item in geodb


	return closestPOIKeyId, nil
}

func (db *GeoDB) Delete(keyId string) error {
	conn := db.pool.Get()
	conn.Do("SELECT", db.dbNum)
	defer conn.Close()

	if _, err := conn.Do("ZREM", db.index, keyId); err != nil {
		return errors.Errorf("Error deleting key %s: %v", keyId, err)
	}

	return nil
}

func (db *GeoDB) Clear() error {
	conn := db.pool.Get()
	conn.Do("SELECT", db.dbNum)
	defer conn.Close()

	if _, err := conn.Do("FLUSHDB"); err != nil {
		return errors.Errorf("Error clearing all key value pairs: %v", err)
	}

	return nil
}
