package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

const Pi = 3.14

func main() {
	// hello world
	fmt.Printf("Hello, world.\n")

	//Math
	fmt.Println(math.Pi)

	//Random number generation
	s3 := rand.NewSource(192891384984)
	r1 := rand.New(s3)
	for i := 0; i < 20; i++ {

		fmt.Print(r1.Intn(100), ",")
	}
	fmt.Print(r1.Intn(100))
	fmt.Println()
	// different function types
	fmt.Println(add(98, 12))
	fmt.Println(multiply(345, 100))
	fmt.Println(swapMultipleReturns("Hello", "World"))
	fmt.Println(splitNamedReturn(90))

	//different variable info
	var c, python, java = "true", "false", 12
	var j int
	k := 54
	fmt.Println(j, c, python, java, k)
	z := cmplx.Sqrt(-5 + 12i)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	var x, y int = 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	var o = uint(f)
	fmt.Println(x, y, o)
	v := 42
	fmt.Printf("v is of type %T\n", v)
	const Truth = true
	fmt.Println("Go rules?", Truth, Pi)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func splitNamedReturn(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func multiply(x, y int) int {
	return x * y
}

func swapMultipleReturns(x, y string) (string, string) {
	return y, x
}

func add(x int, y int) int {
	return x + y
}
