package main

import (
	"fmt"
)

func main() {
	userInput := Input()
	fmt.Println(userInput)
}

func Input() string {
	var userInput string
	fmt.Scan(&userInput)
	return userInput
}
