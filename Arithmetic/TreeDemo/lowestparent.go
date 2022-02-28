package main

import (
	"fmt"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

//完全二叉树
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

	fmt.Println(LowestParent(&node1, &node2, &node5).val)
}

func LowestParent(head *Node, o1 *Node, o2 *Node) *Node {
	if head == nil || head == o1 || head == o2 {
		return head
	}
	left := LowestParent(head.left, o1, o2)
	right := LowestParent(head.right, o1, o2)
	//第一种情况：表示要找的节点在左右树上，意思是两个节点不互为父节点
	if left != nil && right != nil {
		return head
	}
	//第二种情况： 要找的节点，其中一个节点时另一个节点的父节点。
	if left != nil {
		return left
	} else {
		return right
	}
}
