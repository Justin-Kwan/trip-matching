package main

import (
	"fmt"
	"log"

	"order-matching/internal/config"
	"order-matching/internal/transport/websocket"
)

func initWsServer(c *config.WsServerConfig) {
	// fmt.Printf("CONFIG STRUCT: %+v", *cfg)
	fmt.Printf("WS SERVER STRUCT: %+v \n", *c)

	sh := websocket.NewSocketHandler(c)
	log.Printf("Websocket started")
	sh.Serve()
}

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

	initWsServer(&(*cfg).WsServer)
	fmt.Printf("%+v", *cfg)
}
