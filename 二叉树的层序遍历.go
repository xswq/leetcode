package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := make([][]int, 0)
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		size := len(queue)
		level := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			level[i] = node.Val
		}
		ans = append(ans, level)
	}
	return ans
}
