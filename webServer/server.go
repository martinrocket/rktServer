package webServer

import (
	"fmt"
	"net/http"
	"time"
)

func RktWebServer(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Fprintf(w, "%v\n", t.Format("2006-January-02"))
	fmt.Fprintf(w, "Started Web Server....")
}
