package leetcode

import "strings"

//2022-03-15,accept,0ms,2.3MB

func lengthOfLongestSubstring(s string) int {
	res := 0
	str := string("")
	for left, right := 0, 0; right < len(s); right++ {
		if index := strings.IndexByte(str, s[right]); index != -1 {
			left += index + 1
		}
		str = s[left : right+1]
		if len(str) > res {
			res = len(str)
		}
	}
	return res
}
