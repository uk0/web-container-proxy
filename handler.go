package web_container_proxy

// Handler 判断：通过配置文件进行判断
import (
	"net/http"
)

type Handler struct {
	isReverseProxy bool
	isStatic       bool
	uri            string
	server         http.Handler
}
