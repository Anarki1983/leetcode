package leetcode

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
//Memory Usage: 2 MB, less than 88.23% of Go online submissions for Valid Parentheses.

// time complexity: O(N)
// space complexity: O(N)

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]rune, 0, len(s))

	for _, r := range s {
		switch r {
		case '(', '{', '[':
			stack = append(stack, r)
		case ')':
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case '}':
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case ']':
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
