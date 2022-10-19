package leetcode

// time complexity: O(N)
// space complexity: O(1)

func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1

	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum > target {
			j--
		} else {
			i++
		}
	}

	return []int{-1, -1}
}
