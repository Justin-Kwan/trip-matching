package redis

import (
	"github.com/pkg/errors"

	"github.com/gomodule/redigo/redis"
)

type KeyDB struct {
	pool	*redis.Pool
	dbNum int
}

func NewKeyDB(pool *redis.Pool, dbNum int) *KeyDB {
	return &KeyDB{
		pool:  pool,
		dbNum: dbNum,
	}
}

func (k *KeyDB) Select(keyId string) (string, error) {
	conn := k.pool.Get()
  conn.Do("SELECT", k.dbNum)
	defer conn.Close()

	val, err := redis.String(conn.Do("GET", keyId))
	if err != nil {
		return "", errors.Errorf("Error getting value using key '%s': %v", keyId, err)
	}

	return val, nil
}

func (k *KeyDB) Insert(keyId, val string) error {
	conn := k.pool.Get()
  conn.Do("SELECT", k.dbNum)
	defer conn.Close()

	if _, err := conn.Do("SET", keyId, val); err != nil {
		return errors.Errorf("Error setting key '%s': %v", keyId, err)
	}

	return nil
}

func (k *KeyDB) Delete(keyId string) error {
	conn := k.pool.Get()
  conn.Do("SELECT", k.dbNum)
	defer conn.Close()

	if _, err := conn.Do("DEL", keyId); err != nil {
		return errors.Errorf("Error deleting key '%s': %v", keyId, err)
	}

	return nil
}

func (k *KeyDB) Exists(keyId string) (bool, error) {
	conn := k.pool.Get()
  conn.Do("SELECT", k.dbNum)
	defer conn.Close()

	keyExists, err := redis.Bool(conn.Do("EXISTS", keyId))
	if err != nil {
		return false, errors.Errorf("Error checking key '%s' exists: %v", keyId, err)
	}

	return keyExists, nil
}

func (k *KeyDB) CountKeys() (int, error) {
	conn := k.pool.Get()
  conn.Do("SELECT", k.dbNum)
	defer conn.Close()

	val, err := redis.Values(conn.Do("SCAN", nil))
	if err != nil {
		return 0, errors.Errorf("Error counting keys %v:", err)
	}
	keys, _ := redis.Strings(val[1], nil)

	return len(keys), nil
}

func (k *KeyDB) Clear() error {
	conn := k.pool.Get()
  conn.Do("SELECT", k.dbNum)
	defer conn.Close()

	if _, err := conn.Do("FLUSHDB"); err != nil {
		return errors.Errorf("Error clearing all key value pairs: %v", err)
	}

	return nil
}
