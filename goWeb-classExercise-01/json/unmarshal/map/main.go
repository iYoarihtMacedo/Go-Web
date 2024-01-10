package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// string in json format
	stringJSON := `{"name":"John", "age":30, "cars":["Ford", "BMW", "Fiat"]}`

	// transform stringJSON into maps
	// need to use a byte version of original stringJSON
	bytesJSON := []byte(stringJSON)

	var mapJSON map[string]any

	err := json.Unmarshal(bytesJSON, &mapJSON)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v\n", mapJSON)

	// type assertion for access at the values
	value, ok := mapJSON["age"]

	if !ok {
		fmt.Println("age key not found")
		return
	}

	// all numbers taken as floats
	age, ok := value.(float64)
	if !ok {
		fmt.Println("age is not a number")
		return
	}

	fmt.Println("age:", age)

	ageInt := int(age)

	ageInt++

	fmt.Println("age:", ageInt)

}
