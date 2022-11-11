package algorithmproblem

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func swap(nums []int, x, y int) []int {
	nums[x], nums[y] = nums[y], nums[x]
	return nums
}
