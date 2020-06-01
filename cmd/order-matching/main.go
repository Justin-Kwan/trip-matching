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

	// setup redis
	redisPool, err := redis.NewPool(&(*cfg).Redis)
	if err != nil {
		log.Fatalf(err.Error())
	}

	/*rks := */redis.NewKeyDB(redisPool, 0)
	/*rgs := */redis.NewGeoDB(redisPool, 1, "index")

	log.Printf("Redis connection pool initialized...")

	// setup socket server
	sh := websocket.NewSocketHandler(&(*cfg).WsServer)
	sh.Serve()

}
