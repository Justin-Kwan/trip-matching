package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"

	"order-matching/internal/config"
)

// type RedisDb struct {
// 	config *RedisConfig
// 	pool   *redis.Pool
// }

type PoolConfig struct {
	addr            string
	password        string
	connProtocol    string
	idleConnTimeout int
	maxIdleConn     int
	maxActiveConn   int

	// ActiveConn   int    `json:"active_conn"`
	// DB           int    `json:"db"`
	// TobTimeout   string `json:"tob_timeout"`
	// ConnTimeout  string `json:"conn_timeout"`
	// ReadTimeout  string `json:"read_timeout"`
	// WriteTimeout string `json:"write_timeout"`
}

func NewPool(redisCfg *config.RedisConfig) (*redis.Pool, error) {
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
			_, err := conn.Do("PING")
			return err
		},
	}, nil
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
