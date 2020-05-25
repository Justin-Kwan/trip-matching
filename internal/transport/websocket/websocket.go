package websocket

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	// "github.com/pkg/errors"
	
	"order-matching/internal/config"
)

type SocketHandler struct {
	WsServerConfig *WsServerConfig
	upgrader       *websocket.Upgrader
	conn           *websocket.Conn
}

type WsServerConfig struct {
	ReadDeadline int
	ReadTimeout  int
	WriteTimeout int
	MsgSizeLimit int
	Addr         string
	Path         string
}

func (sh *SocketHandler) handleMessage() {
	log.Printf("Client connected!")

	for {
		// conn.SetReadDeadline(time.Now().Add(sh.readDeadline * time.Second))

		msgType, msg, err := sh.conn.ReadMessage()
		if err != nil {
			log.Printf("Client Disconnected!")
			return
		}

		log.Printf("%s sent: %s\n", sh.conn.RemoteAddr(), string(msg))

		if err = sh.conn.WriteMessage(msgType, msg); err != nil {
			log.Printf("Client Disconnected!")
			return
		}
	}
}

func (sh *SocketHandler) handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := sh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	sh.conn = conn
	log.Printf("connection")
	sh.handleMessage()
}

func (sh *SocketHandler) Serve() {
	http.HandleFunc(sh.WsServerConfig.Path, sh.handleConnection)

	svr := &http.Server{
		Addr:         sh.WsServerConfig.Addr,
		ReadTimeout:  time.Duration(sh.WsServerConfig.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(sh.WsServerConfig.WriteTimeout) * time.Second,
	}

	log.Fatal(svr.ListenAndServe())
}

func setConfig(c *config.WsServerConfig) *WsServerConfig {
	return &WsServerConfig{
		ReadDeadline: c.ReadDeadline,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
		MsgSizeLimit: c.MsgSizeLimit,
		Addr:         c.Addr,
		Path:         c.Path,
	}
}

func NewSocketHandler(c *config.WsServerConfig) *SocketHandler {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	return &SocketHandler{
		upgrader:       &upgrader,
		WsServerConfig: setConfig(c),
	}
}
