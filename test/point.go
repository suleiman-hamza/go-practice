package main

import (
	"fmt"
)

func main() {
	letters := "ABC"
	pointer := &letters
	fmt.Println(*pointer)
}
