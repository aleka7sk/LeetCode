package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Val == targetSum && root.Right == nil && root.Left == nil {
		return true
	}
	if hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val) {
		return true
	}

	return false
}
