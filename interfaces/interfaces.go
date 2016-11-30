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

//empty Interface:
type emptyInterface interface{}

//M : This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
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
	//fmt.Println(a.Abs())
	fmt.Println(a)
	a = &f
	fmt.Println(a.Abs())
	fmt.Println("********")
	// this part shows how interface values can be thought of as a tuple of a value and a concrete type
	var i I
	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
	//nil interfaces:
	var interfaceObj I
	describe(interfaceObj)
	//interfaceObj.M()//the panic stuff you see when you uncomment this is what we call a runtime error

	//check empty interface stuff now:
	fmt.Println("****Ëmpty Interfaces:")
	var emp interface{} // this is an empty interface as well
	describeEmpty(emp)
	emp = 42
	describeEmpty(emp)
	emp = "hello"
	describeEmpty(emp)
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

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describeEmpty(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
