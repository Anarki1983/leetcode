package leetcode

// Runtime: 203 ms, faster than 21.31% of Go online submissions for Maximum Subarray.
// Memory Usage: 11.5 MB, less than 5.34% of Go online submissions for Maximum Subarray.

// time complexity: O(N)
// space complexity: O(N)

func maxSubArray(nums []int) int {
	l := len(nums)
	dp := make([]int, l, l)

	switch l {
	case 0:
		return 0
	case 1:
		return nums[0]
	}

	dp[0] = nums[0]
	dp[1] = max(nums[1], dp[0]+nums[1])
	maxSum := max(dp[0], dp[1])
	for i := 2; i < l; i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		if dp[i] > maxSum {
			maxSum = dp[i]
		}
	}

	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
