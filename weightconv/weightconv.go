// Package weightconv performs unit conversions for pounds and kilograms.
package weightconv

import "fmt"

type Pound float64
type Kilogram float64

const lbsToKgFactor = 0.45359237

func (p Pound) String() string    { return fmt.Sprintf("%glb", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }

// PToK converts weight in pounds to kilograms.
func PToK(p Pound) Kilogram { return Kilogram(p * lbsToKgFactor) }

// KToP converts weight in kilograms to pounds.
func KToP(k Kilogram) Pound { return Pound(k / lbsToKgFactor) }
