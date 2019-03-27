package main

import (
	"fmt"
	"log"
	"net/http"
	//"time"

	"github.com/martinrocket/rktServer/webServer"
)

func main() {
	fmt.Println("rktStart web server starting @ http://localhost:8080/...")


	router := http.NewServeMux()
	router.HandleFunc("/intro", webServer.RktWebServer)
	router.HandleFunc("/", webServer.RktStart)



	err := http.ListenAndServe(":8080", router) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
