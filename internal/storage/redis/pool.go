package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"

	"order-matching/internal/config"
)

type PoolConfig struct {
	addr            string
	password        string
	connProtocol    string
	idleConnTimeout int
	maxIdleConn     int
	maxActiveConn   int
}

func NewPool(redisCfg *config.RedisConfig) *redis.Pool {
	cfg := setConfig(redisCfg)

	return &redis.Pool{
		MaxIdle:     cfg.maxIdleConn,
		MaxActive:   cfg.maxActiveConn,
		IdleTimeout: time.Duration(cfg.idleConnTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(cfg.connProtocol, cfg.addr)

			if err != nil {
				return nil, err
			}

			return conn, nil
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}

			_, err := conn.Do("PING")
			return err
		},
	}
}

func setConfig(redisCfg *config.RedisConfig) *PoolConfig {
	return &PoolConfig{
		idleConnTimeout: redisCfg.IdleTimeout,
		maxIdleConn:     redisCfg.MaxIdle,
		maxActiveConn:   redisCfg.MaxActive,
		addr:            redisCfg.Addr,
		password:        redisCfg.Password,
		connProtocol:    redisCfg.ConnProtocol,
	}
}
