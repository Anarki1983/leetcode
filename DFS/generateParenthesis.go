package leetcode

// Runtime: 3 ms, faster than 59.23% of Go online submissions for Generate Parentheses.
// Memory Usage: 2.7 MB, less than 94.16% of Go online submissions for Generate Parentheses.

var res []string

func generateParenthesis(n int) []string {
	// n=1, ()
	// n=2, (()),()()
	// n=3, ((())),(()()),(())(),()(()),()()()

	res = make([]string, 0, 1)
	generateParenthesisDfs(n, 0, 0, "")

	return res
}

func generateParenthesisDfs(n, left, right int, str string) {
	if len(str) == n*2 {
		res = append(res, str)
		return
	}

	if left < n {
		generateParenthesisDfs(n, left+1, right, str+"(")
	}
	if right < left {
		generateParenthesisDfs(n, left, right+1, str+")")
	}
}
