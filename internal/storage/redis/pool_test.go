package redis

import (
	"testing"

	"order-matching/internal/config"

	"github.com/stretchr/testify/assert"
)

var (
	_testRedisCfg *config.RedisConfig
)

type PoolTestConstants struct {
	configFilePath string
	env            string
}

func newPoolTestConstants() *PoolTestConstants {
	return &PoolTestConstants{
		configFilePath: "../../../",
		env:            "test",
	}
}

func setupPoolTests() {
	tc := newPoolTestConstants()

	testCfg, _ := config.NewConfig(tc.configFilePath, tc.env)
	_testRedisCfg = &(*testCfg).Redis
}

func TestNewPool(t *testing.T) {
	setupPoolTests()

	pool, err := NewPool(_testRedisCfg)
	assert.NoError(t, err, "should create redis connection pool without error")
	assert.Equal(t, 500, pool.MaxIdle, "should set max idle connections")
	assert.Equal(t, 1200, pool.MaxActive, "should set max active connections")
}
