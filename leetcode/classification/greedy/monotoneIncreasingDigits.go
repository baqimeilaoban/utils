package greedy

import "strconv"

// monotoneIncreasingDigits 单调递增的数字
/*
当且仅当每个相邻位数上的数字 x 和 y 满足 x <= y 时，我们称这个整数是单调递增的。
给定一个整数 n ，返回 小于或等于 n 的最大数字，且数字呈 单调递增 。
示例1：
输入: n = 10
输出: 9
示例2：
输入: n = 1234
输出: 1234
示例3：
输入: n = 332
输出: 299
*/
// 从前往后会改变已排序的顺序，要从后往前进行排序 strNum[i-1]>strNum[i] => strNum[i-1]--,strNum[i]=9
func monotoneIncreasingDigits(n int) int {
	strNum := []byte(strconv.Itoa(n))
	flag := len(strNum)
	for i := len(strNum) - 1; i > 0; i-- {
		if strNum[i-1] > strNum[i] {
			strNum[i-1]--
			flag = i
		}
	}
	for i := flag; i < len(strNum); i++ {
		strNum[i] = '9'
	}
	res, _ := strconv.Atoi(string(strNum))
	return res
}
