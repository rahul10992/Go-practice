package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// a sender can close a channel to signify that no more values are to be added

// a receiver can test whether a channel has been closed by assigning a second parameter to the receive expression:
// v, ok := <-ch
//only a sender should close the channel. Never the receiver
// closing is necessary just to tell the receiver that no more values are coming
//such as to terminate a range loop!
//The select statement lets a goroutine wait on multiple communication operations
//A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
// a default case is to try a send or receive without a block
func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("--select statement--")
	ch := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		quit <- 0
	}()
	fibonacciSelect(ch, quit)

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
