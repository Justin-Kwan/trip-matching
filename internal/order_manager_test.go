package internal

import (
	"os"
	// "log"
	"testing"

	// "github.com/stretchr/testify/assert"

	"order-matching/internal/config"
	"order-matching/internal/storage/redis"
)

const (
	// test config dependencies
	_env            = "test"
	_configFilePath = "../"
)

var (
	_orderManager *OrderManager

	testOrderBytes1 = `{
  "orderInfo": {
    "location":{
      "lon": 43.45123431,
      "lat": 75.13124123
    },
    "description": "test_order_description1",
    "consumerId": "test_order_consumer_id1",
    "bidPrice": 100.23
    }
  }`
)

func TestMain(m *testing.M) {
	rdb := beforeAll()
	code := m.Run()
	afterAll(rdb)
	os.Exit(code)
}

func beforeAll() *redis.RedisDb {
	testCfg, _ := config.NewConfig(_configFilePath, _env)
	testRedisCfg := &(*testCfg).Redis
	rdb, _ := redis.NewRedisDb(testRedisCfg)
	rdb.Clear()
	_orderManager = NewOrderManager(rdb)
  return rdb
}

func afterAll(rdb *redis.RedisDb) {
  rdb.Clear()
}

func TestAddOrder(t *testing.T) {
	_orderManager.AddOrder(testOrderBytes1)

}
