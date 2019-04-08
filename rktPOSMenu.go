package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//Menu is a struct of Menu Groups and is exported.
type Menu struct {
	MenuName   string `json:"menu_name"`
	Active     bool   `json:"Active"`
	MenuGroups []MenuGroup
}

type MenuGroup struct {
	ID        int    `json:"ID"`
	Name      string `json:"menu_group_name`
	MenuItems []MenuItem
}

type MenuItem struct {
	ID       int    `json:"ID"`
	MenuItem string `json:"MenuItem"`
}

func posMenu() {

	m1 := Menu{"Menu Main", true, []MenuGroup{
		MenuGroup{
			1, "Daily Functions", []MenuItem{
				MenuItem{
					1, "Menu Item 1"}, {
					2, "Menu Item 2"},
			},
		}, {
			2, "Cashier Functions", []MenuItem{
				MenuItem{
					1, "Menu Item 1"}, {
					2, "Menu Item 2"},
			},
		}, {
			2, "Server Functions", []MenuItem{
				MenuItem{
					1, "Menu Item 1"}, {
					2, "Menu Item 2"},
			},
		},
	},
	}

	n, err := json.MarshalIndent(m1, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(n)
	fmt.Println()

	/*for i := range m1.MenuGroup {
		fmt.Printf("\n%v: %v", m1.MenuGroup[i].ID, m1.MenuGroup[i].MenuItem)
	}*/

}
