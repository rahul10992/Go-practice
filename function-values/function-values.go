package main

import (
	"fmt"
	"math"
)

//functions can be passed around and also created as values
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func functionValuesExample() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

}

//function closures:
//A closure is a function that references variables from outside its body.
// the fn may access and assign to the referenced variables.

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func functionClosureExample() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func main() {
	//functionValuesExample()
	functionClosureExample()
}
