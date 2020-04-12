package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast := head
	slow := head
	cut := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		cut = slow
		slow = slow.Next
	}
	cut.Next = nil
	return merge(mergeSort(head), mergeSort(slow))
}

func merge(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
			cur = cur.Next
		} else {
			cur.Next = right
			right = right.Next
			cur = cur.Next
		}
	}
	if left == nil {
		cur.Next = right
	}
	if right == nil {
		cur.Next = left
	}
	return dummy.Next
}
