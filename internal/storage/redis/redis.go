package redis

import (
	"time"

  "github.com/pkg/errors"
	"github.com/gomodule/redigo/redis"

	"order-matching/internal/config"
)

type RedisDb struct {
	config *RedisConfig
	pool   *redis.Pool
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

func NewRedisDb(redisCfg *config.RedisConfig) (*RedisDb, error) {
	rdb := &RedisDb{}
	rdb.config = setConfig(redisCfg)
	rdb.pool = rdb.newConnPool()

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
func (rdb *RedisDb) newConnPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     rdb.config.maxIdle,
		MaxActive:   rdb.config.maxActive,
		IdleTimeout: time.Duration(rdb.config.idleTimeout) * time.Second,

		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(rdb.config.connProtocol, rdb.config.addr)
			if err != nil {
				return nil, errors.Errorf("Error creating redis connection pool: %v", err)
			}
			return conn, err
		},
	}
}

func (rdb *RedisDb) verifyConn() error {
	conn := rdb.pool.Get()
	defer conn.Close()

	if _, err := conn.Do("PING"); err != nil {
		return err
	}
	return nil
}

func (rdb *RedisDb) Select(keyId string) (string, error) {
	conn := rdb.pool.Get()
	defer conn.Close()

	val, err := redis.String(conn.Do("GET", keyId))
	if err != nil {
		return "", errors.Errorf("Error getting value using key '%s': %v", keyId, err)
	}
	return val, nil
}

func (rdb *RedisDb) Insert(keyId string, val string) error {
	conn := rdb.pool.Get()
	defer conn.Close()

	if _, err := conn.Do("SET", keyId, val); err != nil {
		return errors.Errorf("Error setting key '%s': %v", keyId, err)
	}
	return nil
}

func (rdb *RedisDb) Delete(keyId string) error {
	conn := rdb.pool.Get()
	defer conn.Close()

	if _, err := conn.Do("DEL", keyId); err != nil {
		return errors.Errorf("Error deleting key %s: %v", keyId, err)
	}
	return nil
}

func (rdb *RedisDb) Exists(keyId string) bool {
	conn := rdb.pool.Get()
	defer conn.Close()

	exists, _ := redis.Bool(conn.Do("EXISTS", keyId))
	return exists
}

func (rdb *RedisDb) Clear() error {
  conn := rdb.pool.Get()
  defer conn.Close()

  if _, err := redis.Bool(conn.Do("FLUSHDb")); err != nil {
    return errors.Errorf("Error clearing all key value pairs: %v", err)
  }
  return nil
}
