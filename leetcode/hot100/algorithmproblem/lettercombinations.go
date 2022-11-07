package algorithmproblem

// LetterCombinations 电话的字母组合
/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
示例1：
输入：digits = "23" 2:abc 3:def
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
示例2：
输入：digits = ""
输出：[]
示例3：
输入：digits = "2"
输出：["a","b","c"]
*/
func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return combinations
	}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {
	// 若是传入的索引与字符串长度一致时
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		// 获取当前索引数字
		digit := string(digits[index])
		// 获取当前数字代表的字符串
		letters := numMap[digit]
		letterCount := len(letters)
		for i := 0; i < letterCount; i++ {
			// digits 原始字符串 index+1 下一个需要索引的数字
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}

var (
	numMap = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	combinations []string
)
