package main

import (
	// "fmt"
	"go-practice/cf"
	// "go-practice/tempconv"
	"os"
)

func main() {
	Arguement := os.Args[1:]
	// c := tempconv.Celcius(90)
	// f := tempconv.CToF(c)
	// fmt.Println(c, "=", f)

	// fmt.Println(tempconv.BoilingC)
	// fmt.Println(tempconv.FreezingC)
	cf.Best(Arguement)
}
