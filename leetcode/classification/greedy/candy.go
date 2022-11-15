package greedy

// candy 分发糖果
/*
n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
你需要按照以下要求，给这些孩子分发糖果：
每个孩子至少分配到 1 个糖果。
相邻两个孩子评分更高的孩子会获得更多的糖果。
请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。
示例1：
输入：ratings = [1,0,2]
输出：5
解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。
示例2：
输入：ratings = [1,2,2]
输出：4
解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
     第三个孩子只得到 1 颗糖果，这满足题面中的两个条件
*/
func candy(ratings []int) int {
	var candyList []int
	for i := 0; i < len(ratings); i++ {
		candyList = append(candyList, 1)
	}
	// 先从前往后遍历
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candyList[i] = candyList[i-1] + 1
		}
	}
	// 从后往前遍历，最后一个同学没有靠右
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candyList[i] = max(candyList[i], candyList[i+1]+1)
		}
	}
	res := 0
	for i := 0; i < len(candyList); i++ {
		res += candyList[i]
	}
	return res
}
