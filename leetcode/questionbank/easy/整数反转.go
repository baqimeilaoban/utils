package easy

import (
	"math"
	"strconv"
)

// reverse 整数反转
/*
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
假设环境不允许存储 64 位整数（有符号或无符号）。
示例1：
输入：x = 123
输出：321
示例2：
输入：x = -123
输出：-321
示例3：
输入：x = 120
输出：21
示例4：
输入：x = 0
输出：0
*/
func reverse(x int) int {
	flag := false
	if x < 0 {
		flag = true
		x = x * -1
	}
	var res string
	for x > 0 {
		res += strconv.Itoa(x % 10)
		x = x / 10
	}
	num, _ := strconv.Atoi(res)
	if flag {
		num = num * -1
	}
	if num < math.MinInt32 || num > math.MaxInt32 {
		return 0
	}
	return num
}
