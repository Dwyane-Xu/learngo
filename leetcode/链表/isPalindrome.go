package main

// 234. 回文链表

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	var pre, cur *ListNode
	for fast != nil && fast.Next != nil {
		cur = slow

		slow = slow.Next
		fast = fast.Next.Next

		cur.Next = pre
		pre = cur
	}

	if fast != nil {
		slow = slow.Next
	}

	for pre != nil {
		if pre.Val != slow.Val {
			return false
		}
		pre = pre.Next
		slow = slow.Next
	}

	return true

	// if fast != nil {
	// 	slow = slow.Next
	// }
	//
	// var pre, next *ListNode
	// for slow != nil {
	// 	next = slow.Next
	// 	slow.Next = pre
	// 	pre = slow
	// 	slow = next
	// }
	//
	// first := head
	// for pre != nil {
	// 	if pre.Val != first.Val {
	// 		return false
	// 	}
	// 	first = first.Next
	// 	pre = pre.Next
	// }
	//
	// return true
}

func main() {

}
