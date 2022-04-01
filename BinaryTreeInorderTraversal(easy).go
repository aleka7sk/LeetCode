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

var arr = []int{}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	inOrder(root)
	mak := []int{}
	for _, elem := range arr {
		mak = append(mak, elem)
	}
	arr = nil
	return mak
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	arr = append(arr, root.Val)
	inOrder(root.Right)
}
