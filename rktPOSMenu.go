package main

import (
	"encoding/json"
	"log"
	"os"
)

//Menu is a struct of Menu Groups and is exported.
type Menu struct {
	MenuName   string `json:"menu_name"`
	Active     bool   `json:"Active"`
	MenuGroups []MenuGroup
}

// MenuGroup is data model for menu group lists
type MenuGroup struct {
	ID        int    `json:"ID"`
	Name      string `json:"menu_group_name`
	MenuItems []MenuItem
}

// MenuItem is data model for menu item list
type MenuItem struct {
	ID       int    `json:"ID"`
	MenuItem string `json:"MenuItem"`
}

func posMenu() ([]byte, []byte) {

	m1 := Menu{"Menu Main", true, []MenuGroup{
		MenuGroup{
			1, "Daily Functions", []MenuItem{
				MenuItem{
					1, "Start of Day"}, {
					2, "Count Safe"}, {
					3, "Assign Drawers"}, {
					4, "Prevous Day Clock Out Report"}, {
					5, "Count Drawer"}, {
					6, "Cash Drop to Safe"}, {
					7, "Close Drawer"}, {
					8, "End of Day"},
			},
		}, {
			2, "Cashier Functions", []MenuItem{
				MenuItem{
					1, "Start Drawer Count"}, {
					2, "End Drawer Count"}, {
					3, "Drop Cash"}, {
					4, "Clock Out"},
			},
		}, {
			3, "Server Functions", []MenuItem{
				MenuItem{
					1, "Enter Tip"}, {
					2, "Side Work List"}, {
					3, "Clock Out"},
			},
		},
	},
	}

	n, err := json.Marshal(m1)
	nIndent, err := json.MarshalIndent(m1, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(n)
	return n, nIndent

}
