package main

func deleteNode2(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}

	pre := head
	for pre.Next != nil {
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
			break
		}
		pre = pre.Next
	}

	return head
}

func main() {

}
