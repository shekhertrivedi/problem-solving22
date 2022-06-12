package main

import "fmt"

func main() {
	fmt.Println("main..!!")
	// l1 => 2 -> 4 -> 3 -> nil
	n13 := &Node{Value: 3, Next: nil}
	n12 := &Node{Value: 4, Next: n13}
	n11 := &Node{Value: 2, Next: n12}

	// l2 => 5 -> 6 -> 4 -> nil
	n23 := &Node{Value: 4, Next: nil}
	n22 := &Node{Value: 6, Next: n23}
	n21 := &Node{Value: 5, Next: n22}

	fmt.Println(addLinkedList(n11, n21))

}

type Node struct {
	Value int
	Next  *Node
}

func addLinkedList(l1, l2 *Node) int {
	first := 0
	second := 0
	multiplier := 1
	for l1 != nil {
		first = first + (multiplier * l1.Value)
		l1 = l1.Next
		multiplier = multiplier * 10
	}
	multiplier = 1
	for l2 != nil {
		second = second + (multiplier * l2.Value)
		l2 = l2.Next
		multiplier = multiplier * 10
	}

	return first + second
}
