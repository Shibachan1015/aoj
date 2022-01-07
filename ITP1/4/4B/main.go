package main

import (
	"fmt"
	"math"
)
func main() {
	var r float64
	fmt.Scanf("%f\n", &r)
	fmt.Printf("%f %f\n", math.Pi*float64(r*r), math.Pi*float64(r*2))
}
