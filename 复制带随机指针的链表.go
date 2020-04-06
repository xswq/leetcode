package leetcode

type ListNode struct {
	Val    int
	Next   *ListNode
	Random *ListNode
}

func copyRandomList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		temp := &ListNode{}
		temp.Val = cur.Val
		temp.Next = cur.Next

		cur.Next = temp
		cur = temp.Next
	}
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}

		cur = cur.Next.Next
	}

	clone := head.Next
	cloneCur := clone
	cur = clone.Next
	head.Next = cur
	for cur != nil {
		cloneCur.Next = cur.Next
		cloneCur = cloneCur.Next
		cur.Next = cloneCur.Next
		cur = cur.Next
	}
	return clone
}
