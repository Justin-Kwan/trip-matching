package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	// "order-matching/internal/config"
)

type RedisGeoStore struct {
	pool     *redis.Pool
	dbNum    int
	index    string
	distUnit string
}

func NewGeoStore(pool *redis.Pool, dbNum int, index string) *RedisGeoStore {
	return &RedisGeoStore{
		pool:  pool,
		dbNum: dbNum,
		index: index,
	}
}

func (rgs *RedisGeoStore) Insert(keyId string, coord map[string]float64) error {
	conn := rgs.pool.Get()
	conn.Do("SELECT", rgs.dbNum)
	defer conn.Close()

	_, err := conn.Do("GEOADD", rgs.index, coord["lon"], coord["lat"], keyId)
	if err != nil {
		return errors.Errorf("Error adding POI with key '%s': %v", keyId, err)
	}

	return nil
}

func (rgs *RedisGeoStore) Select(keyId string) (map[string]float64, error) {
	conn := rgs.pool.Get()
	conn.Do("SELECT", rgs.dbNum)
	defer conn.Close()

	val, err := redis.Positions(conn.Do("GEOPOS", rgs.index, keyId))
	if err != nil || val[0] == nil {
		return nil, errors.Errorf("Error selecting POI with key '%s'", keyId)
	}

	coord := map[string]float64{
		"lon": val[0][0],
		"lat": val[0][1],
	}

	return coord, nil
}

// func (rgs *RedisGeoStore) SelectAllInRadius(coords map[string]float64, radius float64) ? {
// 	conn := rgs.pool.Get()
// 	defer conn.Close()
//
// 	_, err := conn.Do(
// 		"GEORADIUS",
// 		rgs.index,
// 		coords["lon"],
// 		coords["lat"],
// 		radius,
// 		"km",
// 		"WITHCOORD")
//
// 	if err != nil {
// 		return "", errors.Errorf("Error adding POI with key '%s': %v", keyId, err)
// 	}
// 	return val, nil
// }

func (rgs *RedisGeoStore) Delete(keyId string) error {
	conn := rgs.pool.Get()
	conn.Do("SELECT", rgs.dbNum)
	defer conn.Close()

	if _, err := conn.Do("ZREM", rgs.index, keyId); err != nil {
		return errors.Errorf("Error deleting key %s: %v", keyId, err)
	}

	return nil
}

func (rgs *RedisGeoStore) Clear() error {
	conn := rgs.pool.Get()
	conn.Do("SELECT", rgs.dbNum)
	defer conn.Close()

	if _, err := conn.Do("FLUSHDB"); err != nil {
		return errors.Errorf("Error clearing all key value pairs: %v", err)
	}

	return nil
}
