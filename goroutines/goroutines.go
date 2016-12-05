package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	go say("hello")
	say("world")

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

//goroutines are lightweight threads managed at go runtime
// just use the keyword 'go' before a function call et voila
//goroutines run in the same address space so the mem must be synchronized
//The sync package provides useful primitives, although you won't need them much in Go as there are other primitives

//mainly there are channels
//Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
//the data flows in the direction of the arrow
//and also, sends and receives are blocked until the other side is ready.
