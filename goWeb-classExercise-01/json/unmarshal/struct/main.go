package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string  `json.name`
	Age    int     `json.age`
	Height float64 `json.height`
}

func main() {

	// string in json format
	stringJSON := `{"name":"John", "age":30, "height":2.5}`

	// transform stringJSON into maps
	// need to use a byte version of original stringJSON
	bytesJSON := []byte(stringJSON)

	var p Person
	if err := json.Unmarshal(bytesJSON, &p); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(p)
}
