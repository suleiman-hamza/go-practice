// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import (
	"fmt"
)

type Celcius float64
type Fahrenhiet float64

const (
	AbsoluteZeroC Celcius = -273.15
	BoilingC      Celcius = 100
	FreezingC     Celcius = 0
)

// This is a method and it belongs to a Type, it operates on the type Celcius
func (c Celcius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenhiet) String() string {
	return fmt.Sprintf("%g°F", f)
}
