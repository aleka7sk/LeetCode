package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	var next *ListNode
	for head.Next != nil {
		next = head.Next
		if prev != nil {
			head.Next = prev
		} else {
			head.Next = nil
		}
		prev = head
		head = next
	}
	head.Next = prev
	return head
}
