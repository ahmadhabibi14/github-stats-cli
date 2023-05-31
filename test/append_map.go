package main

import "fmt"

func main() {
	dataArray := []map[string]int{
		{"JavaScript": 24410},
		{"Go": 2946742},
		{"HTML": 1201},
		{"CSS": 59},
		{"JavaScript": 1000},
	}

	dataLang := []map[string]int{}
	for _, data := range dataArray {
		for key, value := range data {
			langMap := make(map[string]int)
			langMap[key] = value
			dataLang = append(dataLang, langMap)
		}
	}

	fmt.Println(dataLang)
}
