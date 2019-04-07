package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//Menu is a struct of Menu Groups and is exported.
type Menu struct {
	MenuName  string `json:"menu_name"`
	Active    bool
	MenuGroup []MenuItems `json:"menu_group"`
}

type MenuItems struct {
	ID       int
	MenuItem string
}

func posMenu() {

	m1 := &Menu{
		MenuName: "Menu April 2019",
		Active:   true,
		MenuGroup: []MenuItems{
			MenuItems{
				ID:       1,
				MenuItem: "Start of Day",
			}, {
				ID:       2,
				MenuItem: "End of Day",
			},
		},
	}

	m2 := &Menu{}
	m2.MenuName = "Menu May 2019"
	m2.Active = false

	n, err := json.MarshalIndent(m1, "", "  ")
	o, err := json.MarshalIndent(m2, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(n)
	fmt.Println()
	os.Stdout.Write(o)
	for i := range m1.MenuGroup {
		fmt.Printf("\n%v:", m1.MenuGroup[i].MenuItem)
	}

	fmt.Println()

}
