package main

import "fmt"

type ErrNegativeSqrt struct {
	What float64
}

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot sqrt negative number : %v", e.What)
}

func run(x float64) error {
	return &ErrNegativeSqrt{
		x,
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

//Get square root using newton's method zn+1 = (zn - (((zn^2)-x)/(2*zn)))
func Sqrt(x float64) (float64, error) {
	if x <= 0 {
		return x, run(x)
	}
	z := 2.0
	for i := 0; i < 10; i++ {
		z = computeValue(z, x)
	}
	return z, nil
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
