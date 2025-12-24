package main

import (
	"fmt"
)

func main() {
	const boilingF, freezingF = 32, 212
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
