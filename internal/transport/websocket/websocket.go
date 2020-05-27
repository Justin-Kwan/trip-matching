package websocket

import (
	"fmt"
	"log"
	"net/http"
	"time"

	// "github.com/pkg/errors"
	"github.com/gorilla/websocket"

	"order-matching/internal/config"
)

type SocketHandler struct {
	config   *WsServerConfig
	upgrader *websocket.Upgrader
	client   *websocket.Conn
}

type WsServerConfig struct {
	ReadDeadline int
	ReadTimeout  int
	WriteTimeout int
	MsgSizeLimit int
	Addr         string
	Path         string
}

// TODO: inject handler with services/db needed!

func NewSocketHandler(wsCfg *config.WsServerConfig) *SocketHandler {
	upgrader := websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	return &SocketHandler{
		config:   newConfig(wsCfg),
		upgrader: &upgrader,
		client:   nil,
	}
}

func newConfig(wsCfg *config.WsServerConfig) *WsServerConfig {
	return &WsServerConfig{
		ReadDeadline: wsCfg.ReadDeadline,
		ReadTimeout:  wsCfg.ReadTimeout,
		WriteTimeout: wsCfg.WriteTimeout,
		MsgSizeLimit: wsCfg.MsgSizeLimit,
		Addr:         wsCfg.Addr,
		Path:         wsCfg.Path,
	}
}

func (sh *SocketHandler) Serve() {
	http.HandleFunc(sh.config.Path, sh.handleConnection)

	svr := &http.Server{
		Addr:         sh.config.Addr,
		ReadTimeout:  time.Duration(sh.config.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(sh.config.WriteTimeout) * time.Second,
	}

	log.Printf("\nWebsocket started")
	log.Fatal(svr.ListenAndServe())
}

func (sh *SocketHandler) handleConnection(w http.ResponseWriter, r *http.Request) {
	// call controller?? -> auth before upgrading the connection
	// pass a service's interface and respond in callback

	conn, err := sh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Fprintf(w, "Error upgrading connection")
		log.Fatalf(err.Error())
	}

	sh.client = conn
	log.Printf("connection")
	sh.handleMessage()
}

func (sh *SocketHandler) handleMessage() {
	log.Printf("Client connected!")

	for {
		// conn.SetReadDeadline(time.Now().Add(sh.readDeadline * time.Second))

		msgType, msg, err := sh.client.ReadMessage()
		if err != nil {
			// ??
			// return errors.Errorf("Invalid environment flag: %s", env)
			log.Printf("Client Disconnected!")
			return
		}

		log.Printf("%s sent: %s\n", sh.client.RemoteAddr(), string(msg))

		if err = sh.client.WriteMessage(msgType, msg); err != nil {
			log.Printf("Client Disconnected!")
			return
		}
	}
}
