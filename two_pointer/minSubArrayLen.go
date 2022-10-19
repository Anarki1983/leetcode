package leetcode

import "math"

// time complexity: O(N)
// space complexity: O(1)

func minSubArrayLen(target int, nums []int) int {

	minLength := math.MaxInt
	i := 0
	sum := 0
	for j := 0; j < len(nums); j++ {
		sum += nums[j]

		for sum >= target {
			minLength = min(minLength, j-i+1)
			sum -= nums[i]
			i++
		}
	}

	if minLength == math.MaxInt {
		return 0
	}

	return minLength
}

//func min(x, y int) int {
//	if x < y {
//		return x
//	}
//
//	return y
//}
