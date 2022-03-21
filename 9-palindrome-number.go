package leetcode

// 2022-03-21,accepted,8ms,4.4MB

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	rem := x
	y := 0
	for rem != 0 {
		// 取最后一位
		t := rem % 10
		// 将y*10并加上最后一位
		y = y*10 + t
		// rem丢弃最后一位
		rem = rem / 10
	}
	return y == x
}
