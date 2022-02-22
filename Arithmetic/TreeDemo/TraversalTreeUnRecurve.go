package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func main() {
	node1 := Node{val: 1}
	node2 := Node{val: 2}
	node3 := Node{val: 3}
	node4 := Node{val: 4}
	node5 := Node{val: 5}
	node6 := Node{val: 6}
	node7 := Node{val: 7}
	node1.left = &node2
	node1.right = &node3
	node2.left = &node4
	node2.right = &node5
	node3.left = &node6
	node3.right = &node7

	FirstPrint1(&node1)
}

func FirstPrint1(head *Node) {
	if head == nil {
		return
	}

	sk := stack.Stack{}
	sk.Push(head)
	for sk.Len() != 0 {
		head := sk.Pop().(*Node)
		fmt.Println(head.val)
		if head.right != nil {
			sk.Push(head.right)
		}
		if head.left != nil {
			sk.Push(head.left)
		}
	}
}
