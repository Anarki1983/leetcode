package leetcode

// time complexity: O(N)
// space complexity: O(N)

func addBinary(a string, b string) string {
	// bit[i] = a[i]+b[i]
	//   1 1 0
	//     1 1
	// 1 0 0 1

	// if bit[i] == 0 or 1, carry = 0
	// if bit[i] == 2, bit = 0, carry = 1

	// 注意slice長度 & carry 最終是否已處理

	// return reverse bit

	// *hack, do not reverse

	aSlice := []rune(a)
	bSlice := []rune(b)

	al := len(a)
	bl := len(b)

	resStr := make([]rune, 0, al)
	bit := '0'
	carry := '0'
	for i, j := al-1, bl-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		x := '0'
		y := '0'
		if i >= 0 {
			x = aSlice[i]
		}
		if j >= 0 {
			y = bSlice[j]
		}

		bit, carry = bitAdd(x, y, carry)

		resStr = append(resStr, bit)
	}

	if carry == '1' {
		resStr = append(resStr, carry)
	}

	return string(reverse(resStr))
}

func reverse(s []rune) []rune {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func bitAdd(a rune, b rune, carry rune) (rune, rune) {
	// 0 + 0 + carry
	if a == '0' && b == '0' {
		return carry, '0' // dig, carry
	}

	// 1 + 1
	if a == '1' && b == '1' {
		// 1 + 1 + carry
		if carry == '1' {
			return '1', '1' // dig, carry
		}

		// 1 + 1
		return '0', '1' // dig, carry
	}

	// 1 + 0 + 1
	if carry == '1' {
		return '0', '1'
	}

	// 1 + 0
	return '1', '0'
}
