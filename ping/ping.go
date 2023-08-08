// ping.go

package ping

import (
	"fmt"
	"log"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming request from %s for path %s", r.RemoteAddr, r.URL.Path)
	fmt.Fprint(w, "Pong!!!")
}
