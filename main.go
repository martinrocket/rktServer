package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

var version = appInfo{"rktPOS", 0.1, 2019, runtime.GOOS}

func main() {
	clear()
	fmt.Println("Main function running...")
	fmt.Printf("version call = %v %v %v %v\n", version.name, version.version, version.year, version.os)
	ServerLog("version call = " + string(version.name) + " " + fmt.Sprintf("%f", version.version) + " " + strconv.Itoa(version.year) + " " + version.os)
	routes()

}

type appInfo struct {
	name    string
	version float64
	year    int
	os      string
}

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func ServerLog(l string) {
	f, err := os.OpenFile("server_log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Printf(l + "\n")
}
