package main

import (
	"fmt"

	. "github.com/baqimeilaoban/utils/leetcode/hot100/algorithmproblem"
)

func main() {
	//generateAddTwoNumsCondition()
	//generateLengthOfLongestSubstringCondition()
	//generateLongestPalindromeCondition()
	//generateMaxArea()
	//generateThreeSums()
	//generateMaxDepth() // 二叉树最大深度
	//generateMinDepth() // 二叉树最小深度
	//generateLetterCombinations()
	//generateRemoveNthFromEnd()
	//generateGenerateParenthesis()
	//generateMergeKLists()
	//generateNextPermutation() // 下一个排列
	//generateLongestValidParentheses() // 最长有效括号
	//generateSearch() // 寻找螺旋数组
	//generateSearchRange() // 在排序数组中查找元素的第一个和最后一个位置
	//generateCombinationSum() // 组合总和
	//generateTrap() // 接雨水
	//generateRotate() // 旋转图像
	generateGroupAnagrams() // 字母异位分词
}

func generateAddTwoNumsCondition() {
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}

	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}
	res := AddTwoNums(l1, l2)
	fmt.Println(res)
}

func generateLengthOfLongestSubstringCondition() {
	fmt.Println(LengthOfLongestSubstring("pwwkew"))
}

func generateLongestPalindromeCondition() {
	fmt.Println(LongestPalindrome("babad"))
}

func generateMaxArea() {
	fmt.Println(MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func generateThreeSums() {
	fmt.Println(ThreeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func generateMaxDepth() {
	fmt.Println(MaxDepth(generateTree()))
}

func generateMinDepth() {
	fmt.Println(MinDepth(generateTree()))
}

func generateLetterCombinations() {
	fmt.Println(LetterCombinations("8"))
}

func generateRemoveNthFromEnd() {
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 3}
	l1.Next.Next.Next = &ListNode{Val: 4}
	l1.Next.Next.Next.Next = &ListNode{Val: 5}
	res := RemoveNthFromEnd(l1, 2)
	fmt.Println(res.Val, res.Next.Val, res.Next.Next.Val, res.Next.Next.Next.Val)
}

func generateGenerateParenthesis() {
	fmt.Println(GenerateParenthesis(3))
}

func generateMergeKLists() {
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 5}

	l2 := &ListNode{Val: 1}
	l2.Next = &ListNode{Val: 3}
	l2.Next.Next = &ListNode{Val: 4}

	l3 := &ListNode{Val: 2}
	l3.Next = &ListNode{Val: 6}
	var lists []*ListNode
	lists = append(lists, l1)
	lists = append(lists, l2)
	lists = append(lists, l3)
	res := MergeKLists(lists)
	fmt.Println(getListVal(res))
}

func getListVal(list *ListNode) []int {
	var res []int
	for list != nil {
		res = append(res, list.Val)
		list = list.Next
	}
	return res
}

func generateTree() *TreeNode {
	return &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
}

func generateNextPermutation() {
	nums := []int{3, 2, 1}
	NextPermutation(nums)
	fmt.Println(nums)
}

func generateLongestValidParentheses() {
	fmt.Println(LongestValidParentheses("(()"))
}

func generateSearch() {
	fmt.Println(Search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}

func generateSearchRange() {
	fmt.Println(SearchRange([]int{5, 7, 7, 8, 8, 10}, 6))
}

func generateCombinationSum() {
	fmt.Println(CombinationSum([]int{2, 3, 5}, 8))
}

func generateTrap() {
	fmt.Println(Trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func generateRotate() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	Rotate(matrix)
	fmt.Println(matrix)
}

func generateGroupAnagrams() {
	fmt.Println(GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
