package websocket

import (
	"log"
	"time"
  "net/http"

	"github.com/gorilla/websocket"
)

type SocketHandler struct {
	upgrader *websocket.Upgrader
	conn     *websocket.Conn
	wsConfig *WsConfig
}

type WsConfig struct {
	ReadDeadline time.Time
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	MsgSizeLimit int64
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

func (sh *SocketHandler) connect(w http.ResponseWriter, r *http.Request) {
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
	http.HandleFunc(sh.wsConfig.Path, sh.connect)

	svr := &http.Server{
		Addr: sh.wsConfig.Addr,
		ReadTimeout: sh.wsConfig.ReadTimeout * time.Second,
		WriteTimeout: sh.wsConfig.WriteTimeout * time.Second,
	}

	/*log.Fatalf(*/svr.ListenAndServe()//)
}

func NewSocketHandler(wsConfig *WsConfig) *SocketHandler {
	upgrader := websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	return &SocketHandler{
		upgrader: &upgrader,
		wsConfig: wsConfig,
	}
}
