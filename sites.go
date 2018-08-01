package web_container_proxy

import "log"

type Sites map[string]Handlers

var sites Sites

func SetupSites(config map[string]map[string]string) {
	newsites := make(Sites)

	for hostname, options := range config {
		log.Println("HostName is  :" + hostname)
		newsites[hostname] = handlersOf(options)
	}

	sites = newsites
}
