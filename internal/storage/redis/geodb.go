package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

type GeoDB struct {
	pool     *redis.Pool
	dbNum    int
	index    string
}

func NewGeoDB(pool *redis.Pool, dbNum int, index string) *GeoDB {
	return &GeoDB{
		pool:  pool,
		dbNum: dbNum,
		index: index,
	}
}

func (g *GeoDB) Insert(keyId string, coord map[string]float64) error {
	conn := g.pool.Get()
	conn.Do("SELECT", g.dbNum)
	defer conn.Close()

	_, err := conn.Do("GEOADD", g.index, coord["lon"], coord["lat"], keyId)
	if err != nil {
		return errors.Errorf("Error adding POI with key '%s': %v", keyId, err)
	}

	return nil
}

func (g *GeoDB) Select(keyId string) (map[string]float64, error) {
	conn := g.pool.Get()
	conn.Do("SELECT", g.dbNum)
	defer conn.Close()

	val, err := redis.Positions(conn.Do("GEOPOS", g.index, keyId))
	if err != nil || val[0] == nil {
		return nil, errors.Errorf("Error selecting POI with key '%s'", keyId)
	}

	coord := map[string]float64{
		"lon": val[0][0],
		"lat": val[0][1],
	}

	return coord, nil
}

// func (g *GeoDB) SelectAllInRadius(coords map[string]float64, radius float64) ? {
// 	conn := g.pool.Get()
// 	defer conn.Close()
//
// 	_, err := conn.Do(
// 		"GEORADIUS",
// 		g.index,
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

func (g *GeoDB) Delete(keyId string) error {
	conn := g.pool.Get()
	conn.Do("SELECT", g.dbNum)
	defer conn.Close()

	if _, err := conn.Do("ZREM", g.index, keyId); err != nil {
		return errors.Errorf("Error deleting key %s: %v", keyId, err)
	}

	return nil
}

func (g *GeoDB) Clear() error {
	conn := g.pool.Get()
	conn.Do("SELECT", g.dbNum)
	defer conn.Close()

	if _, err := conn.Do("FLUSHDB"); err != nil {
		return errors.Errorf("Error clearing all key value pairs: %v", err)
	}

	return nil
}
