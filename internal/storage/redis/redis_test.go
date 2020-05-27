package redis

import (
  "log"
	"testing"

	// "github.com/stretchr/testify/assert"
	// "github.com/gomodule/redigo/redis"

	"order-matching/internal/config"
)

var (
	_configFilePath = "../../../"
	_env            = "test"
)

func testSetup() (*RedisDb, error) {
  testCfg, _ := config.NewConfig(_configFilePath, _env)
  rdb, err := NewRedisDb(&(*testCfg).Redis)
  if err != nil {
    return nil, err
  }
  return rdb, nil
}

func TestNewRedisDb(t *testing.T) {
  rdb, err := testSetup()
  if err != nil {
    log.Fatalf(err.Error())
  }
  rdb.Insert("key", "value")
  rdb.Delete("key")

  // value, err := rdb.Get("key")
  // if err != nil {
  //   log.Fatalf("err")
  // }

  // assert.Equal(t, "value", value, "get")
}
