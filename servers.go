package web_container_proxy

// 核心代码  包含静态服务器（文件以及页面） 反向代理
import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func ReverseProxyServer(uri string) http.Handler {
	dest, _ := url.Parse(addProtocol(uri))
	return httputil.NewSingleHostReverseProxy(dest)
}

func newStaticServer(uri string, hasCustom404 bool, custom404 string) http.Handler {
	log.Println("newStaticServer is 1" + uri)
	return &StaticServer{http.FileServer(http.Dir(GetCurrPath() + uri)), hasCustom404, custom404}
}

func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}

func newSingleFileServer(uri string) http.Handler {
	log.Println("newSingleFileServer is 2" + uri)
	return &SingleFileServer{uri}
}

func addProtocol(url string) string {
	if matches, _ := regexp.MatchString("^\\w+://", url); !matches {
		return fmt.Sprintf("http://%s", url)
	}

	return url
}
