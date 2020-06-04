package main

import (
	"log"

	"order-matching/internal/config"
	"order-matching/internal/storage/redis"
	"order-matching/internal/transport/websocket"
)

const (
	_configFilePath = "../../"
)

func initRedis(cfg *config.Config) error {
	keyDBPool := redis.NewPool(&(*cfg).RedisKeyDB)
	geoDBPool := redis.NewPool(&(*cfg).RedisGeoDB)

	/*keyDB := */ redis.NewKeyDB(keyDBPool)
	/*geoDB := */ redis.NewGeoDB(geoDBPool, "index")
	log.Printf("Redis connection pools initialized...")
	return nil
}

func main() {
	env, err := config.ParseEnvFlag()
	if err != nil {
		log.Fatalf(err.Error())
	}

	cfg, err := config.NewConfig(_configFilePath, env)
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("App config: %+v \n", *cfg)

	initRedis(cfg)

	// setup socket server
	sh := websocket.NewSocketHandler(&(*cfg).WsServer)
	sh.Serve()

}
