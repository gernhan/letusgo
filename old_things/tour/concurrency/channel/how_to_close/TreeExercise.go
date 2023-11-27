package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	left := t.Left
	if left != nil {
		Walk(left, ch)
	}

	ch <- t.Value

	right := t.Right
	if right != nil {
		Walk(right, ch)
	}
}

func Walker(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walker(t1, ch1)

	ch2 := make(chan int)
	go Walker(t2, ch2)

	for i := range ch1  {
		cmp := <-ch2
		if i != cmp {
			return false
		}
		fmt.Printf("Channel1 : %v, Channel2 : %v\n", i, cmp)
	}

	return true
}

func main() {
	fmt.Println(Same(tree.New(2), tree.New(2)))
}
