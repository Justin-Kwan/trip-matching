package redis

import (
	// "log"
	"testing"

	. "github.com/franela/goblin"
	// "github.com/pkg/errors"

	"order-matching/internal/config"
)

func TestGeoStore(t *testing.T) {
  g := Goblin(t)

  env := "test"
  configFilePath := "../../../"
  dbNum := 0

  var rgs *RedisGeoStore

  g.Describe("keystore.go tests", func() {

    g.Before(func() {
      testCfg, _ := config.NewConfig(configFilePath, env)
      testRedisCfg := &(*testCfg).Redis
      rdb, _ := NewRedisDb(testRedisCfg)
      rgs = NewRedisGeoStore(rdb, dbNum)
      rgs.Clear()
    })

    g.AfterEach(func() {

    })

    g.Describe("Insert() Tests", func() {

      g.It("should insert a key with a small value", func() {
        geoQuery := &GeoQuery{
          keyId: "id",
          lon: 99.2,
          lat: 12.21313,
        }

        rgs.Insert(geoQuery)
      })

    })

  })

}
