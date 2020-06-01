package redis

import (
	"github.com/pkg/errors"

	"github.com/gomodule/redigo/redis"
)

type RedisKeyStore struct {
	db    *RedisDb
	dbNum int
}

func NewRedisKeyStore(rdb *RedisDb, dbNum int) *RedisKeyStore {
	return &RedisKeyStore{
		db:    rdb,
		dbNum: dbNum,
	}
}

func (rks *RedisKeyStore) Select(keyId string) (string, error) {
	conn := rks.db.pool.Get()
  conn.Do("SELECT", rks.dbNum)
	defer conn.Close()

	val, err := redis.String(conn.Do("GET", keyId))
	if err != nil {
		return "", errors.Errorf("Error getting value using key '%s': %v", keyId, err)
	}
	return val, nil
}

func (rks *RedisKeyStore) Insert(keyId, val string) error {
	conn := rks.db.pool.Get()
  conn.Do("SELECT", rks.dbNum)
	defer conn.Close()

	if _, err := conn.Do("SET", keyId, val); err != nil {
		return errors.Errorf("Error setting key '%s': %v", keyId, err)
	}
	return nil
}

func (rks *RedisKeyStore) Delete(keyId string) error {
	conn := rks.db.pool.Get()
  conn.Do("SELECT", rks.dbNum)
	defer conn.Close()

	if _, err := conn.Do("DEL", keyId); err != nil {
		return errors.Errorf("Error deleting key %s: %v", keyId, err)
	}
	return nil
}

func (rks *RedisKeyStore) Exists(keyId string) (bool, error) {
	conn := rks.db.pool.Get()
  conn.Do("SELECT", rks.dbNum)
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", keyId))
	if err != nil {
		return errors.Errorf("Error checking key '%s' exists: %v", keyId, err)
	}
	return exists
}

func (rks *RedisKeyStore) CountKeys() (int, error) {
	conn := rks.db.pool.Get()
  conn.Do("SELECT", rks.dbNum)
	defer conn.Close()

	keys := []string{}
	arr, err := redis.Values(conn.Do("SCAN", nil))
	if err != nil {
		return 0, errors.Errorf("Error counting keys %v:", err)
	}

	k, _ := redis.Strings(arr[1], nil)
	keys = append(keys, k...)

	return len(keys), nil
}

func (rks *RedisKeyStore) Clear() error {
	conn := rks.db.pool.Get()
  conn.Do("SELECT", rks.dbNum)
	defer conn.Close()

	// why bool?
	if _, err := redis.Bool(conn.Do("FLUSHDB")); err != nil {
		return errors.Errorf("Error clearing all key value pairs: %v", err)
	}
	return nil
}
