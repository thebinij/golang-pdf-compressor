// ping.go

package ping

import (
	"fmt"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Pong!!!")
}
