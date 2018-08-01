package web_container_proxy

import (
	"net/http"
)

func OnRequest(w http.ResponseWriter, r *http.Request) {

	server, found := matchingServerOf(r.Host, r.URL.String())

	if found {
		server.ServeHTTP(w, r)
		return
	}

	http.NotFound(w, r)
}
