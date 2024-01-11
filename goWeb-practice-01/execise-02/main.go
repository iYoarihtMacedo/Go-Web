package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name     string `json.name`
	LastName string `json.lastName`
}

func main() {

	// reading a request
	http.HandleFunc("/greetings", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.Write([]byte("Hello World"))
		case "POST":
			body := json.NewDecoder(r.Body)
			var data Person

			if err := body.Decode(&data); err != nil {
				fmt.Println(err)
				return
			}

			w.Write([]byte("Hello " + data.Name + " " + data.LastName))
		}
	})

	// listen and serve
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
		return
	}

}
