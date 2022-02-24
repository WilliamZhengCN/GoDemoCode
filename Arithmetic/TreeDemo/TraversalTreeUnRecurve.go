package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

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
	fmt.Println("----------")
	MidPrint2(&node1)
	fmt.Println("----------")
	LastPrint1(&node1)
	fmt.Println("----------")
}

func FirstPrint1(head *Node) {
	if head == nil {
		return
	}

	sk := stack.Stack{}
	sk.Push(head)
	for sk.Len() != 0 {
		head := sk.Pop().(*Node)
		fmt.Printf("%v ", head.val)
		if head.right != nil {
			sk.Push(head.right)
		}
		if head.left != nil {
			sk.Push(head.left)
		}
	}
}

func MidPrint2(head *Node) {
	if head == nil {
		return
	}
	stack1 := stack.Stack{}
	curNode := head

	for stack1.Len() != 0 || curNode != nil {
		if curNode != nil {
			stack1.Push(curNode)
			curNode = curNode.left
		} else {
			curNode = stack1.Pop().(*Node)
			fmt.Printf("%v ", curNode.val)
			//if head.right != nil {
			curNode = curNode.right
			//}
		}
	}
}

func MidPrint1(head *Node) {
	if head == nil {
		return
	}
	stack1 := stack.Stack{}
	stack1.Push(head)
	curNode := head
	for curNode.left != nil {
		stack1.Push(curNode.left)
		curNode = curNode.left
	}

	for stack1.Len() != 0 {
		node := stack1.Pop().(*Node)
		fmt.Printf("%v ", node.val)
		if node.right != nil {
			stack1.Push(node.right)
			node = node.right
			for node.left != nil {
				stack1.Push(node.left)
				node = node.left
			}
		}
	}
}

func LastPrint1(node *Node) {
	if node == nil {
		return
	}
	stack1 := stack.Stack{}
	collectStack := stack.Stack{}

	stack1.Push(node)
	for stack1.Len() != 0 {
		collectNode := stack1.Pop().(*Node)
		collectStack.Push(collectNode)
		if collectNode.left != nil {
			stack1.Push(collectNode.left)
		}
		if collectNode.right != nil {
			stack1.Push(collectNode.right)
		}
	}

	for collectStack.Len() != 0 {
		node1 := collectStack.Pop().(*Node)
		fmt.Printf("%v ", node1.val)
	}
}
