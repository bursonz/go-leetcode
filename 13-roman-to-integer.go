package leetcode

//2022-03-21,accepted,4ms,2.7MB

func romanToInt(s string) int {
	result := 0

	for i, pre, now := 0, 0, 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			now = 1
		case 'V':
			now = 5
		case 'X':
			now = 10
		case 'L':
			now = 50
		case 'C':
			now = 100
		case 'D':
			now = 500
		case 'M':
			now = 1000
		}
		if pre < now {
			result += now - 2*pre
		} else {
			result += now
		}
		pre = now
	}
	return result
}
