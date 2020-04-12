package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func inorderTraversal(root *TreeNode) []int {
//	var ans []int
//	helper(root, &ans)
//	return ans
//}
//
//func helper(root *TreeNode, ans *[]int) {
//	if root == nil {
//		return
//	}
//	helper(root.Left, ans)
//	*ans = append(*ans, root.Val)
//	helper(root.Right, ans)
//}

// 迭代

const (
	white = 0
	gray  = 1
)

type Node struct {
	Root  *TreeNode
	Color int
}

func inorderTraversal(root *TreeNode) []int {
	var ans []int
	stack := []*Node{&Node{Root: root}}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Root == nil {
			continue
		}
		if node.Color == white {
			stack = append(stack, &Node{Root: node.Root.Right})
			stack = append(stack, &Node{Root: node.Root, Color: gray})
			stack = append(stack, &Node{Root: node.Root.Left})
		} else if node.Color == gray {
			ans = append(ans, node.Root.Val)
		}
	}
	return ans
}
