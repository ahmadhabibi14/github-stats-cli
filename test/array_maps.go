package main

import (
	"fmt"
)

func main() {
	dataArray := []map[string]int{
		{"JavaScript": 24410, "Go": 2946742},
		{"HTML": 1201},
		{"CSS": 59},
		{"JavaScript": 1000},
	}

	for _, data := range dataArray {
		for key, value := range data {
			fmt.Printf("Key: %s, Value: %d\n", key, value)
		}
	}
}
