package handler

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/v2fly/BrowserBridge/handler/websocketadp"
	"github.com/v2fly/BrowserBridge/proto"
)

func (hs HTTPHandle) ServeClient(rw http.ResponseWriter, r *http.Request) {
	if hs.link.bridgemux == nil {
		return
	}
	upg := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}
	conn, err := upg.Upgrade(rw, r, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	wsconn := websocketadp.NewWsAdp(conn)
	stream, err := hs.link.bridgemux.Open()
	if err != nil {
		fmt.Println(err.Error())
		hs.link.bridgemux = nil
	}
	var req proto.WebsocketConnectionRequest
	req.Destination = hs.link.RemoteAddr
	req.DestinationSize = uint32(len(hs.link.RemoteAddr))

	proto.WriteRequest(stream, &req)

	go io.Copy(stream, wsconn)
	io.Copy(wsconn, stream)
	stream.Close()

}

func (hs HTTPHandle) Dial(remoteaddr string) (io.ReadWriteCloser, error) {
	if hs.link.bridgemux == nil {
		return nil, errors.New("link is not connected, please connect your browser to the address")
	}
	stream, err := hs.link.bridgemux.Open()
	if err != nil {
		fmt.Println(err.Error())
		hs.link.bridgemux = nil
	}
	var req proto.WebsocketConnectionRequest
	req.Destination = remoteaddr
	req.DestinationSize = uint32(len(remoteaddr))

	proto.WriteRequest(stream, &req)

	return stream, nil
}

func (hs HTTPHandle) Dial2(remoteaddr, protocol string) (io.ReadWriteCloser, error) {
	if hs.link.bridgemux == nil {
		return nil, errors.New("link is not connected, please connect your browser to the address")
	}
	stream, err := hs.link.bridgemux.Open()
	if err != nil {
		fmt.Println(err.Error())
		hs.link.bridgemux = nil
	}
	var req proto.WebsocketConnectionRequest
	req.Destination = remoteaddr
	req.DestinationSize = uint32(len(remoteaddr))
	req.ProtocolString = protocol
	req.ProtocolStringSize = uint32(len(protocol))

	proto.WriteRequest(stream, &req)

	return stream, nil
}
