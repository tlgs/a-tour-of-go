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

	ch <- t.Value
	Walk(t.Right, ch)
	Walk(t.Left, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	m1 := make(map[int]bool)
	m2 := make(map[int]bool)

	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		m1[<-ch1] = true
		m2[<-ch2] = true
	}

	if len(m1) != len(m2) {
		return false
	}
	for k := range m1 {
		_, ok := m2[k]
		if !ok {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
