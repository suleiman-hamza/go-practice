package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celcius) Fahrenhiet { return Fahrenhiet(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenhiet) Celcius { return Celcius((f - 32) * 5 / 9) }
