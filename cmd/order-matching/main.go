package main

import (
	"log"
  "fmt"

	"order-matching/internal/config"
  "order-matching/internal/transport/websocket"
)

func main() {
  cfg, err := config.NewConfig()
  if err != nil {
    log.Fatalf(err.Error())
  }

  fmt.Printf("%+v", *cfg)
  log.Printf("Websocket started")
  sh := websocket.NewSocketHandler(*cfg.WsConfig)
  //sh.Serve()
}
