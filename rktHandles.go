package main

import (
 "fmt"
  "net/http"
  "time"
)



func handleAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "you got the api")
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "you got the about")
}

func handleCareers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "you got the Careers")
}

func handleRoot(w http.ResponseWriter, r *http.Request) {

	t := time.Now()
	fmt.Fprintf(w, "%v\n", t.Format("2006-January-02"))
	fmt.Fprintf(w, "Started Web Server....")


}

func handleIntro(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, `
    <head>
      <h2>My Header</h2>

    </head>
    <body>

        <button>Hello there</button>

    </body>

    `)
}
