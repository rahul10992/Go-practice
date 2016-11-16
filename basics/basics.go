package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("****Pointers****")
	i, j := 42, 2701
	p := &i // p = address of i
	fmt.Println(p)
	fmt.Println(*p) // print value referenced by p
	*p = 21         //value referenced by p = 21
	fmt.Println(i)

	p = &j       // p = address of j
	*p = *p / 37 // value referenced by p = value referenced by p / 37 i.e. j = j/37
	fmt.Println(j)

	fmt.Println("****Structs****")
	v := Vertex{1, 2}
	fmt.Println(v)
	v.x = 98
	fmt.Println(v)

	//struct pointers
	ptr := &v   // ptr = address of v
	ptr.x = 1e9 // or (*ptr).x which is more accurate but this notation is too cumbersome
	fmt.Println(v)
	fmt.Println(avp)
	fmt.Println(avp.x)
	fmt.Println(avp.y)

	fmt.Println("****Arrays****")
	var a [10]int
	for j := 0; j < 10; j++ {
		a[j] = j * 2
	}
	u := [4]string{"hi", "I", "am", "xyz"}
	fmt.Println(u)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4] //slices are dynamically sized, flexible views into the elements of an array; in this context, the []int is unnecessary
	fmt.Println(s)
	// slices do not store any data of their own.
	// They only provide views into an array are pretty much references to them.
	//So changes in slices mean changes in the array as well
	first := u[0:2]
	last := u[1:3]
	fmt.Println(first, last)
	first[0] = "Hello!!"
	fmt.Println(first, last)
	fmt.Println(u)

	//slice literals -> like array literals without having to mention the size
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(slice1)

	slice2 := []bool{true, false, true, true, false, true}
	fmt.Println(slice2)

	slice3 := []struct { //** note this way of initializing a slice of structs
		i int
		b bool
	}{
		{2, false},
		{3, true},
		{4, true},
		{5, false},
		{76, false},
		{13, false},
	}
	fmt.Println(slice3)

	//for slices, the default bounds start from 0 for lower and the upper bound is at the length of the slice
	var slice4 = []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	fmt.Println(slice4)
	slice4 = slice4[1:6]
	fmt.Println("slice4[1:6]=", slice4)
	slice4 = slice4[:4]
	fmt.Println("slice4[:4]=", slice4)
	slice4 = slice4[1:]
	fmt.Println("slice4[1:]=", slice4)
	//length of the slice is the number of elements it contains
	//capacity is the size of the underlying array counting from the first element of the slice
	slice5 := []int{9, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice5)
	slice5 = slice5[:0]
	fmt.Println("slice5[:0]=", slice5)
	slice5 = slice5[:4]
	fmt.Println("slice5[:4]=", slice5)
	slice5 = slice5[2:]
	fmt.Println("slice5[2:]=", slice5)
	fmt.Println("length: ", len(slice5))
	fmt.Println("Capacity: ", cap(slice5))
	//nil slices are those with zero value
	var nilSlice []int
	fmt.Println(nilSlice, len(nilSlice), cap(nilSlice))
	if nilSlice == nil {
		fmt.Println("nil!")
	}
	//slices can be made with the make() func. It allocates a zeroed array and returns a slice that refers to that array
	slice6 := make([]int, 5)    //length of this slice is 5
	slice7 := make([]int, 0, 5) //length of slice = 0 and capacity = 5
	fmt.Println("slice6", slice6)
	fmt.Println("slice7", slice7)
	//slices can contain any type. Including other slices
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	// the append new elements to the slice
	slice8 := make([]int, 0)
	fmt.Println("slice8", slice8)
	slice8 = append(slice8, 0)
	fmt.Println("slice8", slice8)
	slice8 = append(slice8, 1)
	fmt.Println("slice8", slice8)
	slice8 = append(slice8, 2, 3, 4, 5, 6)
	fmt.Println("slice8", slice8)
	//the 'range' form for loop iterates over a slice or map. Sort of like iterators
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, r := range pow {
		fmt.Printf("%d = %d, ", i, r)
	}
	fmt.Println()
	//other alternatives here
	for i := range pow {
		fmt.Printf("%d, ", pow[i])
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	fmt.Println()

}

var ( //struct literals -> you denote a newly allocated struct value by listing the values of its fields
	v1  = Vertex{1, 2}  // has type Vertex
	v2  = Vertex{x: 1}  // Y:0 is implicit
	v3  = Vertex{}      // X:0 and Y:0
	avp = &Vertex{1, 2} // has type *Vertex
)

type Vertex struct {
	x int
	y int
}
