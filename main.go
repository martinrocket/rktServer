package main

import (
	"fmt"
	"log"
	"net/http"

)

func main() {
	fmt.Println("rktStart web server starting @ http://localhost:8080/...")

	r := http.NewServeMux()
	r.HandleFunc("/api/", handleAPI)
	r.HandleFunc("/about", handleAbout)
	r.HandleFunc("/careers", handleCareers)
	r.HandleFunc("/intro", handleIntro)
	r.HandleFunc("/", handleRoot)

	err := http.ListenAndServe(":8080", r) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
