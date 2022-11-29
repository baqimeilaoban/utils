package easy

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type ListNode struct {
	Val  int
	Next *ListNode
}
