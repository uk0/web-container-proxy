package web_container_proxy

//   处理host 以及 URl
import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func matchingHandlerOf(url, hostname string, handlers Handlers) (result http.Handler, found bool) {

	if handlers == nil {
		return nil, false
	}

	for pattern, handler := range handlers {

		if pattern == "*" {
			continue
		}

		if len(url) >= len(pattern) && url[0:len(pattern)] == pattern {
			found = true
			result = http.StripPrefix(pattern, handler.server)
		}
	}

	if handler, hasDefaultHandler := handlers["*"]; !found && hasDefaultHandler {
		result = handler.server
		found = true
	}

	return result, found
}

func matchingServerOf(host, url string) (result http.Handler, found bool) {
	log.Println(" url:" + url + "host:" + host)
	hostname := hostnameOf(host)
	wildcard := wildcardOf(hostname)

	result, found = matchingHandlerOf(url, hostname, sites[hostname])

	if !found {
		if _, hasWildcard := sites[wildcard]; hasWildcard {
			result, found = matchingHandlerOf(url, hostname, sites[wildcard])
		} else {
			log.Println("[INFO] Handler is found")
		}
	}

	if wildcardSite, hasWildcardSite := sites["*"]; !found && hasWildcardSite {
		result, found = matchingHandlerOf(url, hostname, wildcardSite)
	} else if !found {
		log.Println("[INFO] Handler is found")
	} else {
		log.Println("[INFO] Handler is null")
	}

	return result, found
}

func hostnameOf(host string) string {
	hostname := strings.Split(host, ":")[0]

	if len(hostname) > 4 && hostname[0:4] == "www." {
		hostname = hostname[4:]
	}

	return hostname
}

func wildcardOf(hostname string) string {
	parts := strings.Split(hostname, ".")

	if len(parts) < 3 {
		return fmt.Sprintf("*.%s", hostname)
	}

	parts[0] = "*"
	return strings.Join(parts, ".")

}
