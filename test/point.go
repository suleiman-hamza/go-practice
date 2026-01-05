package main

import (
	"fmt"
)

func main() {
	email := "hamzasu"

	if length := getLength(email); length > 1 {
		fmt.Printf("Email is Valid")
	} else {
		fmt.Printf("Email is not Valid")
	}
}

func getLength(str string) int {
	return len(str)
}
