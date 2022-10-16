package leetcode

// Runtime: 10 ms, faster than 31.49% of Go online submissions for Combination Sum.
// Memory Usage: 3.9 MB, less than 36.74% of Go online submissions for Combination Sum.

// time complexity: O(N^2)
// space complexity: O(1)

func maxArea(height []int) int {
	// two pointer
	// i: start ->, j: <- end
	// hack, skip lower height

	start, end := 0, len(height)-1
	maxA := min(height[start], height[end])
	leftMaxHeight := height[start]
	for i := start; i < end-1; i++ {
		if height[i] < leftMaxHeight {
			continue
		} else {
			leftMaxHeight = height[i]
		}

		rightMaxHeight := height[end]
		for j := end; i < j; j-- {
			if height[j] < rightMaxHeight {
				continue
			} else {
				rightMaxHeight = height[j]
			}

			length := j - i

			area := length * min(height[i], height[j])
			if area > maxA {
				maxA = area
			}
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
