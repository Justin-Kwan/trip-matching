package redis

import (
	"testing"

	. "github.com/franela/goblin"

	// "order-matching/internal/config"
)

func TestRedis(t *testing.T) {
	g := Goblin(t)

	// env := "test"
	// configFilePath := "../../../"

	// var rdb *RedisDb

	g.Describe("redis.go tests", func() {

		// g.Before(func() {
		// 	testCfg, _ := config.NewConfig(configFilePath, env)
		// 	testRedisCfg := &(*testCfg).Redis
		// 	rdb, _ = NewRedisDb(testRedisCfg)
		// })

		// todo: fix memory error
		g.Describe("verifyConnection() Tests", func() {
			// err := rdb.verifyConn()
			// g.Assert(err).Equal(nil)
		})

	})
}

// func TestInsertPOI(t *testing.T) {
// 	rdb.InsertPOI("keascsacascsay", 4.555, 9.13123)
// 	lat, lon, err := rdb.SelectPOI("keascsacascsay")
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}
// 	t.Logf("lon: %v, lat: %v", lon, lat)
// }
