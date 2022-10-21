package leetcode

func titleToNumber(columnTitle string) int {
	runes := []rune(columnTitle)

	ret := 0
	multiple := 1
	for i := len(runes) - 1; i >= 0; i-- {
		ret += int(runes[i]-'A'+1) * multiple
		multiple *= 26
	}

	return ret
}
