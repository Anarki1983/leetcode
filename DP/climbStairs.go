package leetcode

// time complexity: O(N)
// space complexity: O(N)

func climbStairs(n int) int {
	// n = 1, 1
	// 1

	// n = 2, 2
	// 1 + 1
	// 2

	// n = 3, 3
	// 1 + 1 + 1
	// 1 + 2
	// 2 + 1

	// n = 4, 5
	// 1 + 1 + 1 + 1
	// 1 + 2 + 1
	// 1 + 1 + 2
	// 2 + 1 + 1
	// 2 + 2

	// f(n) = f(n-1) + f(n-2)
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	f := make([]int, n, n)
	f[0] = 1
	f[1] = 2

	for i := 2; i < n; i++ {
		f[i] = f[i-1] + f[i-2]
	}

	return f[n-1]
}
