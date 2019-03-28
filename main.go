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

	r := http.NewServeMux()
	r.HandleFunc("/api/", handleAPI)
	r.HandleFunc("/about", handleAbout)
	r.HandleFunc("/careers", handleCareers)
	r.HandleFunc("/intro", webServer.RktWebServer)
	r.HandleFunc("/", webServer.RktStart)

	err := http.ListenAndServe(":8080", r) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func myRoutes() {

}

func handleAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "you got the api")
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "you got the about")
}

func handleCareers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "you got the Careers")
}
