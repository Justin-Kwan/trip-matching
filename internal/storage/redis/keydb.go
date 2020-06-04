package redis

import (
	"github.com/pkg/errors"

	"github.com/gomodule/redigo/redis"
)

type KeyDB struct {
	pool   *redis.Pool
}

func NewKeyDB(pool *redis.Pool) *KeyDB {
	return &KeyDB{
		pool: pool,
	}
}

func (db *KeyDB) Insert(keyId, res string) error {
	conn := db.pool.Get()
	defer conn.Close()

	res, err := redis.String(conn.Do("SET", keyId, res))
	if err != nil {
		return errors.Errorf("Error inserting key '%s'", keyId)
	}

	return nil
}

func (db *KeyDB) Select(keyId string) (string, error) {
	conn := db.pool.Get()
	defer conn.Close()

	res, err := redis.String(conn.Do("GET", keyId))
	if err != nil {
		return "", errors.Errorf("Error getting value using key '%s'", keyId)
	}

	return res, nil
}

func (db *KeyDB) Delete(keyId string) error {
	conn := db.pool.Get()
	defer conn.Close()

	res, err := redis.Bool(conn.Do("DEL", keyId))

	keyNotFound := res == false
	if err != nil || keyNotFound {
		return errors.Errorf("Error deleting key '%s'", keyId)
	}

	return nil
}

func (db *KeyDB) Exists(keyId string) (bool, error) {
	conn := db.pool.Get()
	defer conn.Close()

	keyExists, err := redis.Bool(conn.Do("EXISTS", keyId))
	if err != nil {
		return false, errors.Errorf("Error checking key '%s' exists: %v", keyId, err)
	}

	return keyExists, nil
}

func (db *KeyDB) CountKeys() (int, error) {
	conn := db.pool.Get()
	defer conn.Close()

	res, err := redis.Values(conn.Do("SCAN", nil))
	if err != nil {
		return 0, errors.Errorf("Error counting keys %v:", err)
	}

	keys, _ := redis.Strings(res[1], nil)
	return len(keys), nil
}

func (db *KeyDB) Clear() error {
	conn := db.pool.Get()
	defer conn.Close()

	if _, err := conn.Do("FLUSHDB"); err != nil {
		return errors.Errorf("Error clearing all key value pairs: %v", err)
	}

	return nil
}
