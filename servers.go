package web_container_proxy

// 核心代码  包含静态服务器（文件以及页面） 反向代理
import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
)

func ReverseProxyServer(uri string) http.Handler {
	dest, _ := url.Parse(addProtocol(uri))
	return httputil.NewSingleHostReverseProxy(dest)
}

func newStaticServer(uri string, hasCustom404 bool, custom404 string) http.Handler {
	return &StaticServer{http.FileServer(http.Dir(uri)), hasCustom404, custom404}
}

func newSingleFileServer(uri string) http.Handler {
	return &SingleFileServer{uri}
}

func addProtocol(url string) string {
	if matches, _ := regexp.MatchString("^\\w+://", url); !matches {
		return fmt.Sprintf("http://%s", url)
	}

	return url
}
