package leetcode

func canWinNim(n int) bool {
	// 第一把開始抓的時候不是4的倍數就能贏

	return n%4 != 0
}
