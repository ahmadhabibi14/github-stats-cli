package main

import (
	"fmt"
	"sort"
)

func main() {
	// Sample language sizes
	languageSizes := map[string]int{
		"JavaScript": 24410,
		"HTML":       1201,
		"CSS":        59,
		"Python":     8000,
		"Go":         3500,
		"Java":       10000,
		"C++":        5000,
		"C#":         6000,
		"Ruby":       2500,
		"Rust":       2000,
	}

	// Create the map to store language percentages
	lang_with_percentage := make(map[string]float64)

	// Calculate total size of all languages
	totalSize := 0
	for _, size := range languageSizes {
		totalSize += size
	}

	// Determine the number of iterations for the loop
	iterations := len(languageSizes)
	if iterations > 8 {
		iterations = 8
	}

	// Sort the map by values in descending order
	sortedLanguages := sortMapByValueDesc(languageSizes)

	// Calculate and update the language percentages
	for i := 0; i < iterations; i++ {
		language := sortedLanguages[i].Key
		size := sortedLanguages[i].Value
		percentage := float64(size) / float64(totalSize) * 100

		lang_with_percentage[language] = percentage
	}

	// Print the updated map
	fmt.Println(lang_with_percentage)
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
