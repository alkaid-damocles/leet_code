package code

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {

	stack := [][]int{}
	index := 0
	res := []int{}
	for head != nil {
		res = append(res, 0)
		for len(stack) != 0 && stack[len(stack)-1][0] < head.Val {
			res[stack[len(stack)-1][1]] = head.Val
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, []int{head.Val, index})
		index++
		head = head.Next
	}
	return res
}
