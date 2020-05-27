package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"

	"order-matching/internal/config"
)

type RedisDb struct {
	config   *RedisConfig
}

type RedisConfig struct {
	exp          int
	idleTimeout  int
	maxIdle      int
	maxActive    int
	addr         string
	password     string
	connProtocol string
}

type Db interface {
	StartConnPool() error
	Get(keyId string) (string, error)
	// GetAll() ([]string, error)
	Set(keyId string, value string) error
	Delete(keyId string) error
}

var (
  _pool *redis.Pool
)

func NewRedisDb(redisCfg *config.RedisConfig) (*RedisDb, error) {
	rdb := &RedisDb{
		config: setConfig(redisCfg),
	}

  if err := rdb.createConnPool(); err != nil {
    return nil, err
  }
  if err := rdb.verifyConn(); err != nil {
		return nil, err
	}
	return rdb, nil
}

func setConfig(redisCfg *config.RedisConfig) *RedisConfig {
	return &RedisConfig{
		exp:          redisCfg.Exp,
		idleTimeout:  redisCfg.IdleTimeout,
		maxIdle:      redisCfg.MaxIdle,
		maxActive:    redisCfg.MaxActive,
		addr:         redisCfg.Addr,
		password:     redisCfg.Password,
		connProtocol: redisCfg.ConnProtocol,
	}
}

// Sets a redis connection pool to the redis database struct using
// the configuration struct's values.
func (rdb *RedisDb) createConnPool() error {
	_pool = &redis.Pool{
		MaxIdle:     rdb.config.maxIdle,
		MaxActive:   rdb.config.maxActive,
		IdleTimeout: time.Duration(rdb.config.idleTimeout) * time.Second,

		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(rdb.config.connProtocol, rdb.config.addr)
			if err != nil {
				return nil, errors.Errorf("Error creating redis connection pool %v", err)
			}
			return conn, err
		},
	}

  return nil
}

func (rdb *RedisDb) verifyConn() error {
	conn := _pool.Get()
	defer conn.Close()

	if _, err := conn.Do("PING"); err != nil {
		return err
	}
	return nil
}

// Gets and returns a value based on it's key in the redis store.
func (rdb *RedisDb) Get(keyId string) (string, error) {
	conn := _pool.Get()
	defer conn.Close()

	val, err := redis.String(conn.Do("GET", keyId))
	if err != nil {
		return "", errors.Errorf("Error getting value using key %s: %v", keyId, err)
	}
	return val, nil
}

func (rdb *RedisDb) Set(keyId string, value string) error {
	conn := _pool.Get()
	defer conn.Close()

	if _, err := conn.Do("SET", keyId, value); err != nil {
		return errors.Errorf("Error setting key %s: %v", keyId, err)
	}
	return nil
}

func (rdb *RedisDb) Delete(keyId string) error {
	conn := _pool.Get()
	defer conn.Close()

	if _, err := conn.Do("DEL", keyId); err != nil {
		return errors.Errorf("Error deleting key %s: %v", keyId, err)
	}
	return nil
}

// Gets all values and returns them in an string array
// func (rdb *RedisDb) GetAll() ([]string, error) {
// 	conn := db.pool.Get()
// 	defer conn.Close()
//
// 	iter := 0
// 	values := []string{}
//
// 	for {
// 		arr, err := redis.MultiBulk(conn.Do("SCAN", iter))
// 		if err != nil {
// 			return values, fmt.Errorf("error retrieving '%s' keys", pattern)
// 		}
//
// 		iter, _ = redis.Int(arr[0], nil)
// 		k, _ := redis.Strings(arr[1], nil)
// 		keys = append(keys, k...)
//
// 		if iter == 0 {
// 			break
// 		}
// 	}
//
// 	return keys, nil
// }
