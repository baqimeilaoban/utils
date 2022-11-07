package algorithmproblem

// LongestValidParentheses 最长有效括号
/*
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
示例1：
输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"
示例2：
输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"
示例3：
输入：s = ""
输出：0
*/
func LongestValidParentheses(s string) int {
	// 先定义一个返回值
	maxRes := 0
	var stack []int
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxRes = parenthesesMax(maxRes, i-stack[len(stack)-1])
			}
		}
	}
	return maxRes
}

func parenthesesMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
