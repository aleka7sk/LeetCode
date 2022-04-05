package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	x := minDepth(root.Left)
	y := minDepth(root.Right)

	if x != 0 && y == 0 {
		return 1 + x
	} else if x == 0 && y != 0 {
		return 1 + y
	}
	if x > y {
		return 1 + y
	} else {
		return 1 + x
	}
}
