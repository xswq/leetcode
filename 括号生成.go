package leetcode

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	path := make([]byte, 0)
	ans := make([]string, 0)
	dfs(n, n, path, &ans)
	return ans
}

func dfs(left, right int, path []byte, ans *[]string) {
	if left == 0 && right == 0 {
		temp := make([]byte, len(path))
		copy(temp, path)
		*ans = append(*ans, string(temp))
		return
	}

	if left > 0 {
		path = append(path, '(')
		dfs(left-1, right, path, ans)
		path = path[:len(path)-1]
	}

	if right > left {
		path = append(path, ')')
		dfs(left, right-1, path, ans)
	}
}
