package web_container_proxy

import (
	"fmt"
	"log"
	"net/http"
)

func Listen(port int) {

	http.HandleFunc("/", OnRequest)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Println("is error")
	}
}
