package leetcode

//2022-03-15,accept,12ms,4.4MB

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := new(ListNode)
	p := l1
	q := l2
	r := root
	sum := 0
	for {
		//遍历l1
		if p != nil {
			sum += p.Val
			p = p.Next
		}
		//遍历l2
		if q != nil {
			sum += q.Val
			q = q.Next
		}

		//生成结点
		r.Val = sum % 10
		sum /= 10

		//是否还有数据，是否需要进位
		if p == nil && q == nil && sum == 0 {
			break
		}

		r.Next = new(ListNode)
		r = r.Next
	}

	return root
}
