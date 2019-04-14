package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SheetsMenu1 struct {
	Range          string     `json:"range"`
	MajorDimension string     `json:"majorDimension"`
	Values         [][]string `json:"values"`
}

func readMenu() {

	url := "https://sheets.googleapis.com/v4/spreadsheets/1vYfZsswH6R_yREKU4GpzJwy6DoWOr362nTf4ON5eKHs/values/Menu1"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("key", "406751898019-do286h3vuanp5v8i385g6j3sd5g10s2h.apps.googleusercontent.com")
	req.Header.Add("Authorization", "Bearer ya29.GlvrBrNqTnho7E3R1GO7wfI3X1u_LzePREUPxYuHIBc0HEa3h5aGBKLNy0T2aw6USJY3L7MmQ0vQzfgAvMToyVifRMNV5muU00_3ON_ST2OdfiIU9Q30zu9z-F6w")
	req.Header.Add("cache-control", "no-cache")
	//req.Header.Add("Postman-Token", "a6352d16-19ef-417c-846c-00eea47564d2")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	fmt.Println(string(body))

	var dat SheetsMenu1
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat.MajorDimension)
	fmt.Println(dat.Range)
	for _, v := range dat.Values {
		fmt.Println(v)
	}

}
