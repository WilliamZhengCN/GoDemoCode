package main

/*
import (
	"fmt"
	"math"

	"github.com/golang-collections/collections/queue"
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
	//node5 := Node{val: 5}
	//node6 := Node{val: 6}
	//node7 := Node{val: 7}
	node1.left = &node2
	node1.right = &node3
	node2.left = &node4
	//node2.right = &node5
	//node3.left = &node6
	//node3.right = &node7

	PrintWid(&node1)
}

func PrintWid(head *Node) {
	if head != nil {
		mapNode := make(map[*Node]int)
		curLevl := 1
		curNodes := 0
		max := 0
		mapNode[head] = curLevl

		queue := queue.Queue{}
		queue.Enqueue(head)
		for queue.Len() != 0 {
			node := queue.Dequeue().(*Node)
			levl := mapNode[node]
			if levl == curLevl {
				curNodes++
			} else {
				max = int(math.Max(float64(max), float64(curNodes)))
				curLevl++
				curNodes = 1
			}
			fmt.Printf("%v ", node.val)
			if node.left != nil {
				queue.Enqueue(node.left)
				mapNode[node.left] = curLevl + 1
			}
			if node.right != nil {
				queue.Enqueue(node.right)
				mapNode[node.right] = curLevl + 1
			}
		}
		max = int(math.Max(float64(max), float64(curNodes)))
		fmt.Printf("\n%v", max)
	}
}

func PrintWidSearch(head *Node) {
	if head != nil {
		queue := queue.Queue{}
		queue.Enqueue(head)
		for queue.Len() != 0 {
			node := queue.Dequeue().(*Node)
			fmt.Printf("%v ", node.val)
			if node.left != nil {
				queue.Enqueue(node.left)
			}
			if node.right != nil {
				queue.Enqueue(node.right)
			}
		}
	}
}
*/
