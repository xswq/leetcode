package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := make([]int, 0)
	s2 := make([]int, 0)
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}
	carry := 0
	var head *ListNode
	for len(s1) != 0 || len(s2) != 0 || carry != 0 {
		if len(s1) != 0 {
			carry += s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) != 0 {
			carry += s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}
		val := carry % 10
		node := &ListNode{Val: val}
		node.Next = head
		head = node
		carry = carry / 10
	}
	return head
}
