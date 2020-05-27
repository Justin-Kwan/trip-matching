package main

import (
	"log"

	"order-matching/internal/config"
	"order-matching/internal/transport/websocket"
)

const (
	_configFilePath = "../../"
)

func initWsServer(wsCfg *config.WsServerConfig) {
	sh := websocket.NewSocketHandler(wsCfg)
	sh.Serve()
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

	initWsServer(&(*cfg).WsServer)
}
