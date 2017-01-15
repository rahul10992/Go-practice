package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	Inorder(t, ch)
}

func Inorder(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Inorder(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Inorder(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	channel1 := make(chan int, 10)
	channel2 := make(chan int, 10)
	go Walk(t1, channel1)
	go Walk(t2, channel2)
	for i := 0; i < cap(channel1); i++ {
		if <-channel1 != <-channel2 {
			return false
		}
	}
	return true
}

func main() {
	result := Same(tree.New(1), tree.New(2))
	if result {
		fmt.Println("Same")
	} else {
		fmt.Println("Not Same")
	}
}
