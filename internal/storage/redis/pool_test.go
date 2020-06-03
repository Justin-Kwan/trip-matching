package redis

import (
	"testing"

	"order-matching/internal/config"

	"github.com/stretchr/testify/assert"
)

var (
	_keyDBCfg *config.RedisConfig
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

	cfg, _ := config.NewConfig(tc.configFilePath, tc.env)
	_keyDBCfg = &(*cfg).RedisKeyDB
}

func TestNewPool(t *testing.T) {
	setupPoolTests()

	pool := NewPool(_keyDBCfg)
	assert.Equal(t, 500, pool.MaxIdle, "should set max idle connections")
	assert.Equal(t, 1200, pool.MaxActive, "should set max active connections")
}
