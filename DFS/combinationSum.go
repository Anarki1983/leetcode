package leetcode

// Runtime: 3 ms, faster than 84.85% of Go online submissions for Combination Sum.
// Memory Usage: 3.1 MB, less than 83.53% of Go online submissions for Combination Sum.

// time complexity:  與target, candidates有關
// space complexity: O(length_of_longest_combination)

var ret [][]int

func combinationSum(candidates []int, target int) [][]int {
	ret = make([][]int, 0, 0)

	combination := make([]int, 0, 0)
	combinationSumDfs(candidates, target, combination)

	return ret
}

func combinationSumDfs(candidates []int, target int, combination []int) {
	// find maxNum in combination
	sum := 0
	maxNum := -1
	for _, num := range combination {
		sum += num
		if num > maxNum {
			maxNum = num
		}
	}

	for _, num := range candidates {
		// skip num smaller than maxNum, to prevent 2,3,2 combination which is the same with 2,2,3
		if num < maxNum {
			continue
		}

		// match, append to ret
		if sum+num == target {
			copyCombination := copySlice(combination)
			copyCombination = append(copyCombination, num)

			ret = append(ret, copyCombination)
		}
		// can add more candidates, dfs + back tracking
		if sum+num < target {
			combination = append(combination, num)

			combinationSumDfs(candidates, target, combination)

			combination = combination[:len(combination)-1]
		}
	}
}

func copySlice(src []int) []int {
	dst := make([]int, len(src))

	copy(dst, src)

	return dst
}
