package main

import (
	"log"

	"order-matching/internal/config"
	"order-matching/internal/transport/websocket"
)

func initWsServer(wsCfg *config.WsServerConfig) {
	log.Printf("websocket config: %+v \n", *wsCfg)
	sh := websocket.NewSocketHandler(wsCfg)
	log.Printf("\nWebsocket started")
	sh.Serve()
}

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

	initWsServer(&(*cfg).WsServer)
}
