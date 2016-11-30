package main

import "fmt"

func main() {
	var i interface{} = "hello"
	//type assertion provides access to an interface valué's concrete value
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64) // ok becomes the boolean value to denote status making this whole command more of a bool TryGet()
	fmt.Println(f, ok)

	//f = i.(float64) // panic since i does not hold a float
	//fmt.Println(f)
	fmt.Println("*** Type switches: ***")
	do(12)
	do("yhuhuhu")
	do(true)
}

//Typ=e switches: A type switch is a construct that permits several type assertions in series.
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
