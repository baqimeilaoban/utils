package algorithmproblem

// GenerateParenthesis 有效括号的生成
/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
示例1：
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例2：
输入：n = 1
输出：["()"]
*/
func GenerateParenthesis(n int) []string {
	parenthesisRes = make([]string, 0)
	parenthesisDfs(n, 0, 0, "")
	return parenthesisRes
}

func parenthesisDfs(n, lc, rc int, path string) {
	if lc == n && rc == n {
		parenthesisRes = append(parenthesisRes, path)
	} else {
		if lc < n {
			parenthesisDfs(n, lc+1, rc, path+"(")
		}
		if rc < lc {
			parenthesisDfs(n, lc, rc+1, path+")")
		}
	}
}

var parenthesisRes []string
