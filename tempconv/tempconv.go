package tempconv

import (
	"fmt"
)

// Package tempconv performs Celsius and Fahrenheit conversions.

type Celcius float64
type Fahrenhiet float64

const (
	AbsoluteZerpC Celcius = -273.15
	BoilingC      Celcius = 100
	FreezingC     Celcius = 0
)

func (c Celcius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenhiet) String() string {
	return fmt.Sprintf("%g°F", f)
}
