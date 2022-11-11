package algorithmproblem

// Permute 全排列
/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
示例1：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
示例2：
输入：nums = [0,1]
输出：[[0,1],[1,0]]
示例3：
输入：nums = [1]
输出：[[1]]
*/
func Permute(nums []int) [][]int {
	length := len(nums) - 1
	var output []int
	for i := 0; i <= length; i++ {
		output = append(output, nums[i])
	}
	backtrackPermute(length, 0, output)
	return res
}

func backtrackPermute(length, first int, output []int) {
	if first == length {
		tmp := make([]int, length)
		copy(tmp, output)
		res = append(res, output)
		return
	}
	for i := 0; i <= length; i++ {
		swap(output, first, i)
		backtrackPermute(length, first+1, output)
		swap(output, first, i)
	}
}

var res [][]int
