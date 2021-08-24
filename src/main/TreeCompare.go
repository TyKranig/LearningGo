package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	channelTree1, channelTree2 := make(chan int), make(chan int)
	go Walk(t1, channelTree1)
	go Walk(t2, channelTree2)

	for i := 0; i < 10; i++ {
		x, i := <-channelTree1, <-channelTree2
		if x != i {
			return false
		}
	}
	return true
}

func mainTreeCompare() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
}
