package leetcode

// time complexity: O(m*n)
// space complexity: O(m*n)

func uniquePaths(m int, n int) int {
	// init dp
	// 1 1 1 1 1 1
	// 1 0 0 0 0 0
	// 1 0 0 0 0 0
	// 1 0 0 0 0 0

	dp := make([][]int, m, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			}
		}
	}

	// dynamic programming
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
