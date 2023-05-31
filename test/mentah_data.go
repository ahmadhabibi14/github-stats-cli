package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Repository struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// var data map[string]int
// var total int = 0

func main() {
	username := "ahmadhabibi14"
	apiUrl := fmt.Sprintf("https://api.github.com/users/%s/repos", username)
	// Send GET request to fetch user repositories
	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	var repositories []Repository
	err = json.NewDecoder(response.Body).Decode(&repositories)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for index, value := range repositories {
		fmt.Printf("%d. %s\n", index+1, value.Name)
	}
	// fmt.Println(repositories)

	// languageStats := make(map[string]int)
	// totalBytes := 0

	// // Iterate over repositories and collect language statistics
	// for _, repo := range repositories {
	// 	repoUrl := repo.Url + "/languages"
	// 	resp, err := http.Get(repoUrl)
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 		continue
	// 	}
	// 	defer resp.Body.Close()

	// 	var stats RepositoryStats
	// 	err = json.NewDecoder(resp.Body).Decode(&stats)
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 		continue
	// 	}

	// 	// Aggregate language bytes across repositories
	// 	languageStats[stats.Language] += stats.Bytes
	// 	totalBytes += stats.Bytes
	// }

	// // Calculate language percentages
	// languagePercentages := make(map[string]float64)
	// for language, bytes := range languageStats {
	// 	percentage := float64(bytes) / float64(totalBytes) * 100
	// 	languagePercentages[language] = percentage
	// }

	// // Print language percentages
	// for language, percentage := range languagePercentages {
	// 	fmt.Printf("%s: %.2f\n", language, percentage)
	// }

	// resp, err := http.Get("https://api.github.com/repos/ahmadhabibi14/react-quiz-app/languages")
	// if err != nil {
	// 	log.Println("Error ::", err)
	// 	return
	// }
	// defer resp.Body.Close()

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
