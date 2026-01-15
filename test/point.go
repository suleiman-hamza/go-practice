package main

import (
	"fmt"
)

func main() {
	p := new(bool)
	*p = true
	fmt.Print(*p)
}
