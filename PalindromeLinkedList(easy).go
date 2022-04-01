package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func IsPalindrome(head *ListNode) bool {
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		if arr[i] == arr[j] {
			continue
		} else {
			return false
		}
	}
	return true
}
