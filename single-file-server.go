package web_container_proxy

import (
	"net/http"
)

type SingleFileServer struct {
	uri string
}

func (server *SingleFileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, server.uri)
}
