package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var version = appInfo{"rktPOS", 0.1, 2019, runtime.GOOS}

func main() {
	clear()
	posMenu()
	fmt.Println("Main function running...")
	fmt.Printf("version call = %v %v %v %v\n", version.name, version.version, version.year, version.os)
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
