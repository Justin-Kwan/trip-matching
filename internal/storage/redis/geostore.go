package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"

	// "order-matching/internal/config"
)

type RedisGeoStore struct {
	db      *RedisDb
	dbNum    int
	index    string
	distUnit string
}

type GeoQuery struct {
  keyId    string
  lon      float64
  lat      float64
  radius   float64
  distUnit string
}

func NewGeoQuery

func NewRedisGeoStore(rdb *RedisDb, dbNum int) *RedisGeoStore {
	return &RedisGeoStore{
		db:    rdb,
		dbNum: dbNum,
    index: "index name!",
	}
}

func (rgs *RedisGeoStore) Insert(keyId string, location map[string]float64) error {
	conn := rgs.db.pool.Get()
	defer conn.Close()

	_, err := conn.Do(
		"GEOADD",
		"rgs.index",
		location["lon"],
		location["lat"],
		gq.keyId
	)
	
	if err != nil {
		return errors.Errorf("Error adding POI with key '%s': %v", gq.keyId, err)
	}
	return nil
}

func (rgs *RedisGeoStore) Select(keyId string) (map[string]float64, error) {
  conn := rgs.db.pool.Get()
  defer conn.Close()

  coords, err := redis.Positions(conn.Do("GEOPOS", rgs.index, keyId))
  if err != nil {
    return nil, err
  }

	location := map[string]float64{
		"lon": coords[[0][0]],
		"lat": coords[[0][1]],
	}

  return location, nil
}

// func (rgs *RedisGeoStore) SelectAllInRadius(location map[string]float64, radius float64) ? {
// 	conn := rgs.db.pool.Get()
// 	defer conn.Close()
//
// 	_, err := conn.Do("GEORADIUS",
// 		rgs.index,
// 		location["lon"],
// 		location["lat"],
// 		radius,
// 		"km",
// 		"WITHCOORD")
//
// 	if err != nil {
// 		return "", errors.Errorf("Error adding POI with key '%s': %v", keyId, err)
// 	}
// 	return val, nil
// }

func (rgs *RedisGeoStore) Clear() error {
	conn := rgs.db.pool.Get()
  conn.Do("SELECT", rgs.dbNum)
	defer conn.Close()

	if _, err := redis.Bool(conn.Do("FLUSHDB")); err != nil {
		return errors.Errorf("Error clearing all key value pairs: %v", err)
	}
	return nil
}
