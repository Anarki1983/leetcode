package leetcode

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Plus One.
// Memory Usage: 2.1 MB, less than 63.72% of Go online submissions for Plus One.

// time complexity: O(N)
// space complexity: O(1)

func plusOne(digits []int) []int {
	// increase ith, check if carry
	// if reach 0th of digits and carry exist, push front 1 to digits

	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += carry
		if digits[i] == 10 {
			digits[i] = 0
			carry = 1
		} else {
			carry = 0
		}

		if carry == 0 {
			break
		}
	}

	if carry == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
