package leetcode

// Runtime: 10 ms, faster than 31.49% of Go online submissions for Combination Sum.
// Memory Usage: 3.9 MB, less than 36.74% of Go online submissions for Combination Sum.

// time complexity: O(N)
// space complexity: O(1)

func maxArea(height []int) int {
	// two pointer
	// i: start ->, j: <- end

	start, end := 0, len(height)-1
	maxA := min(height[start], height[end])
	for start < end {
		area := (end - start) * min(height[start], height[end])
		maxA = max(maxA, area)

		if height[start] > height[end] {
			end--
		} else {
			start++
		}
	}

	return maxA
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
