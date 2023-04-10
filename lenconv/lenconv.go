// Package lenconv performs length conversions in feet and meters.
package lenconv

import "fmt"

const meterToFeetFactor = 3.281

type Foot float64
type Meter float64

func (f Foot) String() string  { return fmt.Sprintf("%gft", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

// FToM converts feet length to meters.
func FToM(f Foot) Meter { return Meter(f / meterToFeetFactor) }

// MToF converts meter length to feet.
func MToF(m Meter) Foot { return Foot(m * meterToFeetFactor) }
