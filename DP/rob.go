package leetcode

// time complexity: O(N)
// space complexity: O(N)

func rob(nums []int) int {
	// init rob
	l := len(nums)

	// boundary case
	switch l {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	}

	// rob = money at ith rob
	rob := make([]int, l, l)
	// base case
	rob[0] = nums[0]
	rob[1] = max(nums[0], nums[1])

	// rob[i] = find max from (nums[i] + rob[i-2], rob[-1])
	for i := 2; i < l; i++ {
		rob[i] = max(nums[i]+rob[i-2], rob[i-1])
	}

	return rob[l-1]
}

//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//
//	return y
//}
