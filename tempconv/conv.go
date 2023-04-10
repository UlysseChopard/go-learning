package tempconv

// CtoF converts a Celsius temperature to Fahrenheit.
func CtoF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FtoC converts a Fahrenheit tempreature to Celsius.
func FtoC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CtoK converts a Celsius temperature to Kelvin.
func CtoK(c Celsius) Kelvin { return Kelvin(ZeroK + c) }
