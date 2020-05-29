package internal

import (
  "log"
  "testing"

  "github.com/stretchr/testify/assert"

  "order-matching/internal/config"
  "order-matching/internal/storage/redis"
)

func TestMain(m *testing.M) {
	beforeAll()
	code := m.Run()
	afterAll()
	os.Exit(code)
}

func beforeAll() {
	testCfg, _ := config.NewConfig(_configFilePath, _env)
	_testRedisCfg = &(*testCfg).Redis
	_rdb, _ = NewRedisDb(_testRedisCfg)
  _rdb.Clear()
  
}

func afterAll() {
  _rdb.Clear()
}
