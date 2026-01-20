package lengthconv

import "fmt"

type Feet float64
type Meter float64

const (
	Meterconst Meter = 0.3048
	Feetconst  Feet  = 3.280839895
)

func FeetToMeter(ft Feet) Meter {
	value := ft * Feet(Meterconst)
	return Meter(value)
}

func MeterToFeet(m Meter) Feet {
	value := m * Meter(Feetconst)
	return Feet(value)
}

func (f Feet) String() string {
	return fmt.Sprintf("%gft", f)
}

func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}
