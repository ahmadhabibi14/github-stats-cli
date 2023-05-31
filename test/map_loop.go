package main

import (
	"fmt"
)

func main() {
	data := map[string]int{
		"javascript": 24410,
		"HTML":       1201,
		"CSS":        59,
	}

	total := 0
	for _, value := range data {
		total += value
	}

	fmt.Println("Total:", total)
}
