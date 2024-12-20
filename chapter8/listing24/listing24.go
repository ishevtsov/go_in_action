// This sample program demonstrates how to decode a JSON response
// using the json package and NewDecoder function.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Result struct {
	Name     string  `json:"name"`
	Language string  `json:"language"`
	Id       string  `json:"id"`
	Bio      string  `json:"bio"`
	Version  float64 `json:"version"`
}

func main() {
	uri := "https://microsoftedge.github.io/Demos/json-dummy-data/64KB.json"

	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	defer resp.Body.Close()

	var gr []Result
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(gr)
}
