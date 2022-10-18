package leetcode

import "math"

// time complexity: O(N), N = digit number of x
// space complexity: O(1)

func reverse(x int) int {
	sign := true
	if x < 0 {
		sign = false
		x = -x
	}

	rev := 0

	for x > 0 {
		rev = rev*10 + x%10
		x /= 10
	}

	if !sign {
		rev = -rev
	}

	if rev > math.MaxInt32 || rev < math.MinInt32 {
		return 0
	}

	return rev
}
