package cf

import (
	"fmt"
	"go-practice/tempconv"
	"strconv"
)

func Best(Argue []string) {
	for _, arg := range Argue {
		t, _ := strconv.ParseFloat(arg, 64)

		f := tempconv.Fahrenhiet(t)
		c := tempconv.Celcius(t)

		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
