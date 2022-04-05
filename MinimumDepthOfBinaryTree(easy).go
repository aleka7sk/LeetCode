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

	if x == 0 || y == 0 {
		return max(x, y) + 1
	}
	return min(x, y) + 1
}

func max(first, second int) int {
	if first > second {
		return first
	} else {
		return second
	}
}

func min(first, second int) int {
	if first > second {
		return second
	} else {
		return first
	}
}
