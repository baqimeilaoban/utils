package algorithmproblem

// MinDepth 二叉树最小深度
/*
给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明：叶子节点是指没有子节点的节点。
示例1：
输入：root = [3,9,20,null,null,15,7]
输出：2
示例2：
输入：root = [2,null,3,null,4,null,5,null,6]
输出：5
*/
func MinDepth(root *TreeNode) int {
	treeMin = 0
	treeMinDepth(root, 0)
	return treeMin
}

func treeMinDepth(root *TreeNode, deep int) {
	if root == nil {
		return
	}
	deep++
	if root.Left == nil && root.Right == nil {
		if deep < treeMin || treeMin == 0 {
			treeMin = deep
		}
	}
	// 遍历左子树
	treeMinDepth(root.Left, deep)
	// 遍历右子树
	treeMinDepth(root.Right, deep)
}

var treeMin int
