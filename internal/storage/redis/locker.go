package redis

import (
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

type Locker struct {
	pool *redis.Pool
}

func NewLocker(pool *redis.Pool) *Locker {
	return &Locker{
		pool: pool,
	}
}

// Attempts to lock a single key with a ttl. Specific error is
// returned if passed in key is already locked. A generic error
// is returned for all other errors.
// Locking algorithm reference: https://redis.io/commands/set
func (l *Locker) LockKey(keyId string, lockTTL int) error {
	conn := l.pool.Get()
	defer conn.Close()

	// NX sets key if it does not already exist. Returns OK if key was set.
	// Returns nil if key not set (if key is already set).
	// If 0 is returned, key is already set which means the key is
	// locked by another client. Lock is autoreleased by expiry.

	lockKey := string("lock." + keyId)
	log.Printf("aquiring lock on '" + keyId + "'...")

	res, err := conn.Do("SET", lockKey, "token", "NX", "EX", lockTTL)
	if err != nil { // handle generic errors
		return err
	}

	if res == nil {
		log.Printf("key is already locked")
		return errors.Errorf("Error, key is already locked")
	}

	return nil
}

func (l *Locker) UnlockKey(keyId string) error {
	conn := l.pool.Get()
	defer conn.Close()

	lockKey := string("lock." + keyId)
	res, err := redis.Bool(conn.Do("DEL", lockKey))

	keyNotFound := res == false
	if err != nil || keyNotFound {
		return errors.Errorf("Error unlocking key '%s'", keyId)
	}

	return nil
}
