package websocketadp

import (
	"bytes"
	"net"
	"time"

	"github.com/gorilla/websocket"
)

func NewWsAdp(ws *websocket.Conn) *WsAdp {
	return &WsAdp{ws, nil}
}

type WsAdp struct {
	*websocket.Conn
	readbuf *bytes.Reader
}

func (ws *WsAdp) Read(b []byte) (n int, err error) {
	if ws.readbuf != nil {
		n, err = ws.readbuf.Read(b)
		if ws.readbuf.Len() == 0 {
			ws.readbuf = nil
		}
		return
	}
	_, msg, errw := ws.ReadMessage()
	if errw != nil {
		return 0, errw
	}
	ws.readbuf = bytes.NewReader(msg)

	return ws.Read(b)
}

func (ws *WsAdp) Write(b []byte) (n int, err error) {
	return len(b), ws.WriteMessage(websocket.BinaryMessage, b)
}

func (ws *WsAdp) SetDeadline(t time.Time) error {
	return nil
}

func (ws *WsAdp) AsConn() net.Conn {
	return ws
}
