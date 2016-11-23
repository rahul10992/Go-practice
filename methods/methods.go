package main

import (
	"fmt"
	"math"
)

type myFloat float64
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

// we have methods for non struct types as well
func (f myFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// we can also have pointer receivers as you see:
func (v *vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//Scale is func and lets not forget that methods are functions too and can be rewritten as those:
//try to remove the * in vertex and check it out
func Scale(v *vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//while methods with pointer receivers can take either a value or a ptr as the receiver,
//GO interprets them as the same thing

func main() {
	v := vertex{3, 4}
	fmt.Println(v.Abs())
	f := myFloat(-math.Sqrt2)
	fmt.Println(-math.Sqrt2)
	fmt.Println(f.Abs())
	Scale(&v, 10)
	fmt.Println(v.Abs())
	v.Scale(10)
	fmt.Println(v.Abs())
}
