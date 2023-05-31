package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type repository struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

var Repositories []repository
var dataLang []map[string]int
var langMap map[string]int

// var LanguagesPercentage map[string]float64
var total int = 0

func main() {
	username := "ahmadhabibi14"
	apiUrl := fmt.Sprintf("https://api.github.com/users/%s/repos", username)
	// Send GET request to fetch user Repositories
	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading API response:", err)
		return
	}

	err = json.Unmarshal(body, &Repositories)
	if err != nil {
		fmt.Println("Error unwofu:", err)
		return
	}

	for index, value := range Repositories {
		fmt.Printf("%d. %s\n", index+1, value.Name)
		fmt.Println("+===========================================+")

		reposUrl := fmt.Sprintf("%s/languages", value.Url)
		resp, err := http.Get(reposUrl)
		if err != nil {
			log.Println("Error languages::", err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading API response:", err)
			return
		}

		err = json.Unmarshal(body, &langMap)
		if err != nil {
			log.Println("Error unmarshall :: ", err)
			return
		}

		for key, val := range langMap {
			langMap[key] = val
			dataLang = append(dataLang, langMap)
		}

	}

	fmt.Println(dataLang)
	// err = json.NewDecoder(resp.Body).Decode(&data)
	// if err != nil {
	// 	log.Println("Error decode data")
	// 	return
	// }

	// for _, value := range data {
	// 	total += value
	// }

	// fmt.Println(total)
	// fmt.Println(data)
}
