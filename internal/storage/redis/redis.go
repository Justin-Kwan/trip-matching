package redis

import (
	// "log"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"

	"order-matching/internal/config"
)

type RedisDb struct {
	config *RedisConfig
	pool   *redis.Pool
}

type RedisConfig struct {
	idleTimeout  int
	maxIdle      int
	maxActive    int
	addr         string
	password     string
	connProtocol string
}

type KeyQuery struct {
	keyId string
	val   string
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

	// nil for some reason...
	defer conn.Close()

	if _, err := conn.Do("PING"); err != nil {
		return err
	}
	return nil
}
