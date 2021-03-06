package main

// 面试题 02.02. 返回倒数第 k 个节点

func kthToLast(head *ListNode, k int) int {
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow.Val
}

func main() {

}
