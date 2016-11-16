package main

import (
	"fmt"
	"math"
)

func main() {
	no := 1087.0
	res := sqrt(no)
	fmt.Println(res)

	fmt.Println("Difference: ", makePositive(res-math.Sqrt(no)))
}

func makePositive(x float64) float64 {
	if x >= 0 {
		return x
	}
	return (x * -1)
}

//Get square root using newton's method zn+1 = (zn - (((zn^2)-x)/(2*zn)))
func sqrt(x float64) float64 {
	z := 2.0
	for i := 0; i < 10; i++ {
		z = computeValue(z, x)
	}
	return z
}

func computeValue(zn float64, x float64) (zn1 float64) {
	num := numerator(zn, x)
	den := (2 * zn)
	zn1 = zn - (num / den)
	return
}

func numerator(zn float64, x float64) float64 {
	return (zn * zn) - x
}
