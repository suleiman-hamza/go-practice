package main

import (
	"fmt"
	"go-practice/cf"
	"go-practice/tempconv"
)

func main() {
	// c := tempconv.Celcius(90)
	// f := tempconv.CToF(c)
	// fmt.Println(c, "=", f)

	// fmt.Println(tempconv.BoilingC)
	fmt.Println(tempconv.FreezingC)
	cf.Loop()
}
