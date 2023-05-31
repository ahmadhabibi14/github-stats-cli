package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var languages = make(map[string]int)
var item = make(map[string]int)

// var ngab []map[string]int

func main() {
	url := "https://api.github.com/repos/ahmadhabibi14/react-quiz-app/languages"
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error get http response :: ", err)
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&item)
	if err != nil {
		fmt.Println("Error Decode JSON data :: ", err)
		return
	}

	for i := 0; i <= 2; i++ {
		for key, value := range item {
			if val, ok := languages[key]; ok {
				languages[key] += val
			} else {
				languages[key] = value
			}
		}
	}

	// for _, item := range ngab {
	// 	for key, value := range item {
	// 		fmt.Printf("%s : %d\n", key, value)
	// 	}
	// }
	for key, value := range languages {
		fmt.Printf("%s : %d\n", key, value)
	}
}
