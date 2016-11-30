package main

import (
	"fmt"
	"strconv"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// similar to the stringers
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"didn't work. Better luck next time",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	i, err := strconv.Atoi("42dfg") // make this just 42
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
	} else {
		fmt.Println("Converted integer:", i)
	}
}
