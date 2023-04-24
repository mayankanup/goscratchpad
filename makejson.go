package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string = "Anup"
	var address string = "Gachibowli, Hyderabad"
	var personmap map[string]string
	personmap = make(map[string]string)

	// takes input value for name
	//fmt.Print("Enter your name: ")
	//fmt.Scan(&name)

	personmap["name"] = name

	// takes input value for name
	//fmt.Print("Enter your name: ")
	//fmt.Scan(&name)

	personmap["address"] = address

	fmt.Println(personmap)

	barr, err := json.Marshal(personmap)

	if err == nil {
		fmt.Println(string(barr))
	}

}
