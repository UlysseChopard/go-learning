package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/UlysseChopard/go-learning/lenconv"
	"github.com/UlysseChopard/go-learning/tempconv"
	"github.com/UlysseChopard/go-learning/weightconv"
)

var l = flag.Bool("l", false, "only length conversions")
var w = flag.Bool("w", false, "only weight conversions")
var t = flag.Bool("t", false, "only temperature conversions")

func main() {
	flag.Parse()
	displayAll := !(*l || *w || *t)
	for _, arg := range flag.Args() {
		n, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		if displayAll || *t {
			f := tempconv.Fahrenheit(n)
			c := tempconv.Celsius(n)
			k := tempconv.Kelvin(n)
			fmt.Printf("%s = %s, %s = %s, %s = %s\n", f, tempconv.FtoC(f), c, tempconv.CtoF(c), k, tempconv.KtoC(k))
		}
		if displayAll || *w {
			p := weightconv.Pound(n)
			k := weightconv.Kilogram(n)
			fmt.Printf("%s = %s, %s = %s\n", p, weightconv.PToK(p), k, weightconv.KToP(k))
		}
		if displayAll || *l {
			f := lenconv.Foot(n)
			m := lenconv.Meter(n)
			fmt.Printf("%s = %s, %s = %s\n", f, lenconv.FToM(f), m, lenconv.MToF(m))
		}
	}
}
