package leetcode

//2022-03-27,accept,4ms

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for a, t := range nums {
		b, ok := numMap[target-t]
		if ok {
			return []int{a, b}
		}
		numMap[t] = a
	}
	return nil
}
