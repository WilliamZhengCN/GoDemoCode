package main

import (
	"fmt"
)

func main() {
	f := "测试啊"
	r := []rune(f)
	for i := 0; i < len(r); i++ {
		fmt.Print(r[i])
	}
}

func RemSlice(arr []int) {
	arr[0] = 2
}

func Get(arr ...int) {

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	firstNumIndex, secondNumIndex := 0, 0
	firstNum, secondNum := 0, 0
	curLen := 0
	index1, index2 := 0, 0
	if (len1+len2)%2 == 0 {
		secondNumIndex = (len1 + len2) / 2
		if secondNumIndex > 0 {
			firstNumIndex = secondNumIndex - 1
		}
	} else {
		firstNumIndex = (len1 + len2) / 2
		secondNumIndex = firstNumIndex
	}

	for index1 < len1 || index2 < len2 {
		curVal := 0
		if index1 < len1 && index2 < len2 {
			if nums1[index1] < nums2[index2] {
				curVal = nums1[index1]
				index1++
			} else {
				curVal = nums2[index2]
				index2++
			}
			curLen++
		} else if index1 < len1 {
			curVal = nums1[index1]
			index1++
			curLen++
		} else if index2 < len2 {
			curVal = nums2[index2]
			index2++
			curLen++
		}
		if curLen == firstNumIndex+1 {
			firstNum = curVal
		}
		if curLen == secondNumIndex+1 {
			secondNum = curVal
			break
		}
	}

	return ((float64(firstNum)) + float64(secondNum)) / 2
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func testaddTwoNumbers() {
	node1 := ListNode{}
	node1.Val = 1
	node2 := ListNode{}
	node2.Val = 2
	node1.Next = &node2
	node3 := ListNode{}
	node3.Val = 3
	node2.Next = &node3
	node41 := ListNode{}
	node41.Val = 4
	node3.Next = &node41

	cur := ReverNode(&node1)
	fmt.Print(cur)

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

func ReverNode(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}
