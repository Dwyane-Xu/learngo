package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates2(head *ListNode) *ListNode {
	// dummy := new(ListNode)
	// dummy.Next = head
	// pre, cur := dummy, head
	// for cur != nil && cur.Next != nil {
	// 	if cur.Val == cur.Next.Val {
	// 		for cur.Next != nil && cur.Val == cur.Next.Val {
	// 			cur = cur.Next
	// 		}
	// 		pre.Next = cur.Next
	// 		cur = cur.Next
	// 	} else {
	// 		pre = cur
	// 		cur = cur.Next
	// 	}
	// }
	// return dummy.Next

	dummy := new(ListNode)
	dummy.Next = head
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			temp := cur.Next
			for temp.Next != nil && temp.Val == temp.Next.Val {
				temp = temp.Next
			}
			cur.Next = temp.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

func main() {
	head := &ListNode{1,
		&ListNode{2,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{4, nil}}}}}}
	head = deleteDuplicates2(head)
	fmt.Println(head.Val)
}
