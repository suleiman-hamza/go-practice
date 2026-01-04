package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{"Suleiman", 23}
	p := &u
	fmt.Println(*p)
	fmt.Println(u)
}
