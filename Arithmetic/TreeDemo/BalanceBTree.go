package main

/*
import (
	"fmt"
	"math"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

//平衡二叉树
func main() {
	node1 := Node{val: 1}
	node2 := Node{val: 2}
	node3 := Node{val: 3}
	node4 := Node{val: 4}
	node5 := Node{val: 5}
	node6 := Node{val: 6}
	node7 := Node{val: 7}
	node8 := Node{val: 8}
	node9 := Node{val: 9}
	node1.left = &node2
	node1.right = &node3
	node2.left = &node4
	node2.right = &node5
	node3.left = &node6
	node3.right = &node7
	node7.right = &node8
	node8.right = &node9

	fmt.Println(IsBalanceBT(&node1))
}

func IsBalanceBT(head *Node) (bool, int) {
	if head == nil {
		return true, 0
	}
	leftBalance, leftHeight := IsBalanceBT(head.left)
	rightBalance, righeHeight := IsBalanceBT(head.right)
	curHeight := Max(leftHeight, righeHeight) + 1
	curBalance := leftBalance && rightBalance && (math.Abs(float64(leftHeight)-float64(righeHeight)) < 2)
	return curBalance, curHeight
}

func Max(val int, val1 int) int {
	if val > val1 {
		return val
	} else {
		return val1
	}
}
*/
