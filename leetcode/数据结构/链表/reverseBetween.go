package main

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	left := &ListNode{Val: -1}
	left.Next = head
	for i := 0; i < m-1; i++ {
		left = left.Next
	}
	right := left.Next

	mHead := left.Next
	var pre, next *ListNode
	for i := 0; i < n-m+1; i++ {
		next = mHead.Next
		mHead.Next = pre
		pre = mHead
		mHead = next
	}

	right.Next = mHead
	if left.Val != -1 {
		left.Next = pre
		return head
	} else {
		return pre
	}
}

func main() {

}
