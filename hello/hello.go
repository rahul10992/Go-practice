package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
)

const Pi = 3.14

func main() {
	// hello world
	fmt.Printf("Hello, world.\n")
	fmt.Println("Differnet functions and things. Simple Stuff")

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

	//different variables and constants info
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

	// Go does not have a whie.
	uy := 1
	for uy < 10 { // a statement being executed before getting into the for
		uy += uy
		fmt.Println("uy = ", uy)
	}
	//OR
	// sum := 1
	// for sum < 1000 {
	// 	sum += sum
	// }
	// fmt.Println(sum)

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	printOS()
	fmt.Println()
	deferSample()
	consecutiveDefers()

}

func consecutiveDefers() {
	fmt.Println("Counting") //consecutive defers end up everything in a stack and use LIFO to traverse through all the defers once the function returns

	for i := 0; i <= 10; i++ {
		defer fmt.Println(i)
	}

}

func deferSample() {
	defer fmt.Println("World") //defer defers the execution of a statement until its surrounding function returns
	fmt.Println("Hello")
}

func printOS() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { // a statement being executed before getting into the if
		return v
	}
	return lim
}

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

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
