package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	m := map[string]any{
		"name":   "John",
		"age":    30,
		"height": 2.0,
	}

	// transform into json bytes
	bytes, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return

	}

	data := string(bytes)
	fmt.Println(data)

}
