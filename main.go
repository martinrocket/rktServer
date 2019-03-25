package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/martinrocket/rktServer/webServer"
)

func main() {
	fmt.Println("rktStart web server starting @ http://localhost:8080/...")

	http.HandleFunc("/", webServer.RktWebServer)
	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
