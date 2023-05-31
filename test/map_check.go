package main

import (
	"fmt"
)

func main() {
	myMap := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}

	key := "pizza"
	value, ok := myMap[key]
	if ok {
		// Key exists, do something with the value
		fmt.Println("Value:", value)
	} else {
		// Key does not exist
		fmt.Println("Key not found")
	}
}
