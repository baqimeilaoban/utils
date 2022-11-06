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
	generateThreeSums()
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
