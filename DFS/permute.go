package leetcode

var permuteRet [][]int

func permute(nums []int) [][]int {
	permuteRet = make([][]int, 0, 0)
	visited := make([]bool, len(nums), len(nums))
	permutation := make([]int, 0, len(nums))

	permuteDfs(nums, visited, permutation)

	return permuteRet
}

func permuteDfs(nums []int, visited []bool, permutation []int) {
	for i, num := range nums {
		if !visited[i] {
			visited[i] = true

			permutation = append(permutation, num)

			// if all visited, append to ret
			isAllVisited := true
			for _, ok := range visited {
				if !ok {
					isAllVisited = false
					break
				}
			}

			if isAllVisited {
				copyPermutation := copySlice(permutation)
				permuteRet = append(permuteRet, copyPermutation)
			} else {
				permuteDfs(nums, visited, permutation)
			}

			visited[i] = false
			permutation = permutation[:len(permutation)-1]
		}
	}
}

//func copySlice(src []int) []int {
//	dst := make([]int, len(src))
//
//	copy(dst, src)
//
//	return dst
//}
