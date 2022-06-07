package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)

	if t == nil {
		return
	}

	lch := make(chan int)
	go Walk(t.Left, lch)
	for v := range lch {
		ch <- v
	}

	ch <- t.Value

	rch := make(chan int)
	go Walk(t.Right, rch)
	for v := range rch {
		ch <- v
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if !ok1 || !ok2 {
			return ok1 == ok2
		} else if v1 != v2 {
			return false
		}
	}
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Printf("%v ", v)
	}
	fmt.Println()

	fmt.Println("New(1) == New(1):", Same(tree.New(1), tree.New(1)))
	fmt.Println("New(1) == New(2):", Same(tree.New(1), tree.New(2)))
}
