package main

import (
	"fmt"
	"math"
)

type vertex struct {
	X, Y float64
}

//a method is a function with a special receiver argument
// as you can see below the receive is v of type Vertex
//you can use these type of methods as you would methods within a class
//since go does not have any classes
func (v vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := vertex{3, 4}
	fmt.Println(v.Abs())
}
