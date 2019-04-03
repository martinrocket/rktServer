package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type menu struct {
	menuName string `json: "menu name"`
	levelOne
}

type levelOne struct {
	List  []string `json: "list"`
	Close string   `json: "close"`
}

func posMenu() {

	m := menu{
		menuName: "Menu Number One"
		menuName.levelOne.List: {}
	}

	n, err := json.MarshalIndent(m, "  ", "  ")

	os.Stdout.Write(n)
	fmt.Println()

}
