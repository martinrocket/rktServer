package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("rktStart web server starting @ http://localhost:8080/...")
	http.HandleFunc("/", startWebServer)
	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func startWebServer(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Fprintf(w, "%v\n", t.Format("2006-January-02"))
	fmt.Fprintf(w, "Started Web Server....")
}
