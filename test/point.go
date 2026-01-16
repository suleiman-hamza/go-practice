package main

import (
	"fmt"
)

func main() {
	p := new(string)
	q := new(string)
	fmt.Println(*p == *q)
}
