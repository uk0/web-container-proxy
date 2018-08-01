package web_container_proxy

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

type Handlers map[string]*Handler

func handlerOf(uri string, hasCustom404 bool, custom404 string) *Handler {

	handler := &Handler{false, false, uri, nil}
	isStatic := isLocalPath(uri)

	if isStatic && isSingleFile(uri) {
		handler.isStatic = true
		handler.server = newSingleFileServer(uri)
		log.Println("isSingleFile")
	} else if isStatic {
		handler.isStatic = true
		handler.server = newStaticServer(uri, hasCustom404, custom404)
		log.Println("newStaticServer")
	} else {
		handler.isReverseProxy = true
		handler.server = ReverseProxyServer(uri)
		log.Println("ReverseProxyServer")
	}

	return handler
}

func handlersOf(options map[string]string) Handlers {
	handlers := make(Handlers)

	custom404, hasCustom404 := options["*"]

	if hasCustom404 {
		hasCustom404 = isLocalPath(custom404)
	}

	if hasCustom404 {
		custom404 = fmt.Sprintf("%s/index.html", custom404)
	}

	for path, uri := range options {
		handlers[path] = handlerOf(uri, hasCustom404, custom404)
	}

	return handlers
}

func isLocalPath(config string) bool {
	matches, _ := regexp.MatchString("^/", config)
	return matches
}

func isSingleFile(uri string) bool {
	f, err := os.Open(uri)

	if err != nil {
		return false
	}

	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return false
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		return false
	case mode.IsRegular():
		return true
	}

	return false
}
