package leetcode

//Runtime: 9 ms, faster than 81.42% of Go online submissions for Minimum Path Sum.
//Memory Usage: 3.8 MB, less than 100.00% of Go online submissions for Minimum Path Sum.

// time complexity: O(m*n)
// space complexity: O(m*n)

func minPathSum(grid [][]int) int {
	if grid == nil {
		return 0
	}

	m := len(grid)
	n := len(grid[0])

	// calculate edge's sum
	for i := 1; i < m; i++ {
		grid[i][0] = grid[i][0] + grid[i-1][0]
	}

	for j := 1; j < n; j++ {
		grid[0][j] = grid[0][j] + grid[0][j-1]
	}

	// [i][j] = min([i-1][j], [i][j-1]) + [i][j]
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
		}
	}

	return grid[m-1][n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
