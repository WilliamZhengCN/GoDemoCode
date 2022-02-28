package main

import "fmt"

func main() {
	node1 := ListNode{}
	node1.Val = 9
	node2 := ListNode{}
	node2.Val = 9
	node1.Next = &node2
	node3 := ListNode{}
	node3.Val = 9
	node2.Next = &node3
	node41 := ListNode{}
	node41.Val = 9
	node3.Next = &node41

	node4 := ListNode{}
	node4.Val = 9
	node5 := ListNode{}
	node5.Val = 9
	node4.Next = &node5
	node6 := ListNode{}
	node6.Val = 9
	node5.Next = &node6

	head := addTwoNumbers(&node1, &node4)
	fmt.Print(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	isOver := false
	ret := []int{}
	for l1 != nil || l2 != nil {
		val := 0
		if l1 != nil {
			val += l1.Val
		}
		if l2 != nil {
			val += l2.Val
		}
		if isOver {
			val += 1
		}
		if val > 9 {
			val = val % 10
			isOver = true
		} else {
			isOver = false
		}
		ret = append(ret, val)
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if isOver {
		ret = append(ret, 1)
	}

	newNode := new(ListNode)
	newNode.Val = ret[0]
	head := newNode
	for i := 1; i < len(ret); i++ {
		nextNode := ListNode{}
		nextNode.Val = ret[i]
		newNode.Next = &nextNode
		newNode = &nextNode
	}
	fmt.Println(ret)
	return head
}
