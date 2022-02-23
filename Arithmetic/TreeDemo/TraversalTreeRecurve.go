package main

import (
	"fmt"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

func main1() {
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
	FirstPrint(&node1)
	fmt.Println("----------")
	MidPrint(&node1)
	fmt.Println("----------")
	LastPrint(&node1)
	fmt.Println("----------")
}

func FirstPrint(head *Node) {
	if head == nil {
		return
	}
	fmt.Printf("%v ", head.val)
	FirstPrint(head.left)
	FirstPrint(head.right)
}

func MidPrint(head *Node) {
	if head == nil {
		return
	}
	MidPrint(head.left)
	fmt.Printf("%v ", head.val)
	MidPrint(head.right)
}

func LastPrint(head *Node) {
	if head == nil {
		return
	}
	LastPrint(head.left)
	LastPrint(head.right)
	fmt.Printf("%v ", head.val)
}
