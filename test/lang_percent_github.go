package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
)

type repository struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

var Repositories []repository

var languages = make(map[string]int) // all languages with total for whole repositories
var langItem = make(map[string]int)  // this for store temporary data
var LangToFetch = make(map[string]float64)
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
	err = json.NewDecoder(response.Body).Decode(&Repositories)
	if err != nil {
		fmt.Println("Error Decode response :", err)
		return
	}

	fmt.Println("+===========================================+")
	for _, value := range Repositories {
		// fmt.Printf("%d. %s\n", index+1, value.Name)

		// Request languages data from repository
		reposUrl := fmt.Sprintf("%s/languages", value.Url)
		resp, err := http.Get(reposUrl)
		if err != nil {
			log.Println("Error languages::", err)
			return
		}
		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(&langItem)
		if err != nil {
			log.Println("Error decode languages data from repository : ", err)
			return
		}
		for key, value := range langItem {
			if val, ok := languages[key]; ok {
				languages[key] += val
			} else {
				languages[key] = value
			}
		}
	}
	fmt.Println("+===========================================+")
	for _, value := range languages {
		// fmt.Printf("%s : %d\n", key, value)
		total += value
	}
	iterations := len(languages)
	if iterations > 8 {
		iterations = 8
	}
	sortedLanguages := sortMapByValueDesc(languages)
	for i := 0; i < iterations; i++ {
		lang := sortedLanguages[i].Key
		size := sortedLanguages[i].Value
		percentage := float64(size) / float64(total) * 100

		LangToFetch[lang] = percentage
	}
	fmt.Println("+===========================================+")

	for key, value := range LangToFetch {
		perc := fmt.Sprintf("%.2f%", value)
		fmt.Printf("%v : %v\n", key, perc)
	}
	return
}

// Helper function to sort the map by values in descending order
type pair struct {
	Key   string
	Value int
}

type pairList []pair

func (p pairList) Len() int {
	return len(p)
}

func (p pairList) Less(i, j int) bool {
	return p[i].Value > p[j].Value
}

func (p pairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func sortMapByValueDesc(m map[string]int) pairList {
	pairs := make(pairList, len(m))
	i := 0
	for k, v := range m {
		pairs[i] = pair{k, v}
		i++
	}
	sort.Sort(pairs)
	return pairs
}
