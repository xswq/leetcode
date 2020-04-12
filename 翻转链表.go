package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		node := head.Next
		head.Next = prev
		prev = head
		head = node
	}
	return prev
}
