package main

import (
	"fmt"
	"strings"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (node *List[T]) Append(v T) {
	p := node
	for p.next != nil {
		p = p.next
	}
	p.next = &List[T]{nil, v}
}

func (node *List[T]) AppendLeft(v T) {
	node.next = &List[T]{node.next, node.val}
	node.val = v
}

func (node *List[T]) String() string {
	var s []string
	for p := node; p != nil; p = p.next {
		s = append(s, fmt.Sprint(p.val))
	}
	return strings.Join(s, " -> ")
}

func main() {
	li := &List[int]{}
	li.Append(1)
	li.Append(2)
	li.AppendLeft(3)
	fmt.Println(li)

	ls := &List[string]{nil, "You"}
	ls.AppendLeft("Give")
	ls.AppendLeft("Gonna")
	ls.Append("Up")
	ls.AppendLeft("Never")
	fmt.Println(ls)
}
