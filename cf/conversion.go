package cf

import (
	"fmt"
	"go-practice/lengthconv"
	"go-practice/tempconv"
	"strconv"
)

func Best(Argue []string) {
	for _, arg := range Argue {
		t, _ := strconv.ParseFloat(arg, 64)

		f := tempconv.Fahrenhiet(t)
		c := tempconv.Celcius(t)
		fm := lengthconv.Feet(t)
		mf := lengthconv.Meter(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		fmt.Printf("%s = %s, %s = %s\n", mf, lengthconv.MeterToFeet(mf), fm, lengthconv.FeetToMeter(fm))
	}
}
