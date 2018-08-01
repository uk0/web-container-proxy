package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	web_container_proxy "github.com/breakEval13/libs/lib/web_container_proxy" // 项目的根目录
	"github.com/breakEval13/libs/lib/web_container_proxy/json-config"         // 相对路径 = json-config
)

func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}

func main() {
	var port int = 8080
	go func() {
		config := JSONConfig.NewJSONConfig(GetCurrPath()+"/config.json", web_container_proxy.SetupSites)
		config.EnableAutoReload()
	}()

	strPort := strconv.Itoa(port)
	log.Println("作者： 张建新.")
	log.Println("静态服务器|反向代理服务器")
	log.Println("config.json 支持热加载")
	log.Print("server is start on port " + strPort)
	web_container_proxy.Listen(port)
}
