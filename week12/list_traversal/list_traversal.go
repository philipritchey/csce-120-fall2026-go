package main

import (
	"container/list"
	"fmt"
)

func main() {
	RecursiveSliceTraversal([]int{8, 6, 7, 5, 3, 0, 9})
	fmt.Println("==========")
	l := list.New()
	for _, e := range []int{8, 6, 7, 5, 3, 0, 9} {
		l.PushBack(e)
	}
	RecursiveLinkedListTraversal(l.Front())
}

func RecursiveSliceTraversal(slice []int) {
	if len(slice) == 0 {
		return
	}
	fmt.Println("pre:", slice[0]) // stand-in for processing the current element
	RecursiveSliceTraversal(slice[1:])
	fmt.Println("post:", slice[0])
}

func RecursiveLinkedListTraversal(node *list.Element) {
	if node == nil {
		return
	}
	fmt.Println("pre:", node.Value) // stand-in for processing the current element
	RecursiveLinkedListTraversal(node.Next())
	fmt.Println("post:", node.Value)
}
