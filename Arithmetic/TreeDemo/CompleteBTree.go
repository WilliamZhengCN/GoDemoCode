package main

/*
import (
	"fmt"

	"github.com/golang-collections/collections/queue"
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

	fmt.Println(IsCBT(&node1))
}

func IsCBT(head *Node) bool {
	if head != nil {
		isLeaf := false //是否到达叶子节点
		queue := queue.Queue{}
		queue.Enqueue(head)
		for queue.Len() != 0 {
			node := queue.Dequeue().(*Node)
			if node.left == nil && node.right != nil {
				return false
			}
			if isLeaf {
				if node.left != nil || node.right != nil {
					return false
				}
			} else {
				if node.left != nil && node.right == nil {
					isLeaf = true
				}
			}
			if node.left != nil {
				queue.Enqueue(node.left)
			}
			if node.right != nil {
				queue.Enqueue(node.right)
			}
		}
	}
	return true
}
*/
