// ping.go

package ping

import (
	"fmt"
	"log"
	"net/http"
)

var Version = "1.0.0"

func Ping(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming request from %s for path %s", r.RemoteAddr, r.URL.Path)
	response := fmt.Sprintf("Pong!!! Version: %s", Version)
	fmt.Fprint(w, response)
}
