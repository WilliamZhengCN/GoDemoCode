package main

import (
	"fmt"
	"math"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

//搜索二叉树
func main() {
	node1 := Node{val: 5}
	node2 := Node{val: 3}
	node3 := Node{val: 7}
	node4 := Node{val: 2}
	node5 := Node{val: 4}
	node6 := Node{val: 5}
	node7 := Node{val: 8}
	node1.left = &node2
	node1.right = &node3
	node2.left = &node4
	node2.right = &node5
	node3.left = &node6
	node3.right = &node7

	isBst := IsBst(&node1)
	fmt.Println(isBst)
}

var minf = math.MinInt

func IsBst(head *Node) bool {
	if head == nil {
		return true
	}
	isBst := IsBst(head.left)
	if !isBst {
		return false
	} else {
		if head.val < minf {
			return false
		} else {
			minf = head.val
		}
	}
	return IsBst(head.right)
}
