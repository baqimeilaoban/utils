package easy

import "math"

// isValidBST 验证二叉搜索树
/*
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
有效 二叉搜索树定义如下：
1.节点的左子树只包含 小于 当前节点的数。
2.节点的右子树只包含 大于 当前节点的数。
3.所有左子树和右子树自身必须也是二叉搜索树。
示例1：
输入：root = [2,1,3]
输出：true
示例2：
输入：root = [5,1,4,null,null,3,6]
输出：false
解释：根节点的值是 5 ，但是右子节点的值是 4 。
*/
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt, math.MaxInt)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	// 验证左子树
	left := helper(root.Left, lower, root.Val)
	// 验证右子树
	right := helper(root.Right, root.Val, upper)
	return left && right
}
