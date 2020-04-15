package leetcode

func updateMatrix(matrix [][]int) [][]int {
	queue := make([][]int, 0)
	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else if matrix[i][j] == 1 {
				matrix[i][j] = -1
			}
		}
	}
	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	m := len(matrix)
	n := len(matrix[0])
	for len(queue) != 0 {
		point := queue[0]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			x := point[0] + dx[i]
			y := point[1] + dy[i]

			if x >= 0 && x < m && y >= 0 && y < n && matrix[x][y] == -1 {
				matrix[x][y] = matrix[point[0]][point[1]] + 1
				queue = append(queue, []int{x, y})
			}
		}
	}
	return matrix
}
