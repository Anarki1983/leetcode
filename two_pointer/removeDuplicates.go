package leetcode

// time complexity: O(N)
// space complexity: O(1)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	j := 1

	for j < len(nums) {
		if nums[i] == nums[j] {
			nums = append(nums[:j], nums[j+1:]...)
		} else {
			i++
			j++
		}
	}

	return len(nums)
}
