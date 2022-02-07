package handler

import (
	"net/http"

	"github.com/xtaci/smux"
)

type HandleSettings struct {
	ListenAddr string
	RemoteAddr string
}

func Handle(settings *HandleSettings) error {
	var hdl = HTTPHandle{
		link: new(BridgeLink),
	}
	hdl.link.RemoteAddr = settings.RemoteAddr
	return http.ListenAndServe(settings.ListenAddr, hdl)
}

func NewHttpHandle() *HTTPHandle {
	return &HTTPHandle{
		link: new(BridgeLink),
	}
}

type HTTPHandle struct {
	link *BridgeLink
}
type BridgeLink struct {
	bridgemux  *smux.Session
	RemoteAddr string
}

func (hs HTTPHandle) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	requestPath := r.URL.Path[1:]

	switch requestPath {
	case "":
		fallthrough
	case "index.js":
		BridgeResource(rw, r, requestPath)
		break
	case "link":
		hs.ServeBridge(rw, r)
	case "forward":
		hs.ServeClient(rw, r)
	}
}
func BridgeResource(rw http.ResponseWriter, r *http.Request, path string) {
	http.ServeFile(rw, r, path)
}
