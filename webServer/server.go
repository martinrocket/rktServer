package webServer

import (
	"fmt"
	"net/http"
	"time"
)

// RktStart shows web server has started
func RktStart(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Fprintf(w, "%v\n", t.Format("2006-January-02"))
	fmt.Fprintf(w, "Started Web Server....")
}

// RktWebServer shows intro page
func RktWebServer(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Intro....")
}
