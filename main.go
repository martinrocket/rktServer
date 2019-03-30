package main

import (
	"fmt"
	"runtime"
)

var version = appInfo{"rktPOS", 0.1, 2019, ""}

func main() {
	fmt.Println("Main function running...")
	routes()
	version.os = runtime.GOOS
	fmt.Printf("version call = %v\n", version)

}

type appInfo struct {
	name    string
	version float64
	year    int
	os      string
}
