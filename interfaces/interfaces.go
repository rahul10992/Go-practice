package main

import (
	"fmt"
	"math"
)

//Abser : dummy comment to stop that stupid warning popping out
//interfaces are implemented implicitly no extra keywork reqd
type Abser interface {
	Abs() float64
}

// I interface
type I interface {
	M()
}

// T struct of type t
type T struct {
	S string
}

//M : This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat implements Abser
	fmt.Println(a.Abs())
	a = &v // a *Vertex implements Abser
	fmt.Println(a.Abs())
	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a = v
	a = &f

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
