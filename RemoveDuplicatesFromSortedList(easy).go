package leetcode

/**
 * Definition for singly-linked list.
type ListNode struct {
	Val int
    Next *ListNode
}
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	current := head
	for current.Next != nil {
		if current.Val == current.Next.Val {
			if current.Next.Next == nil {
				current.Next = nil
			} else {
				current.Next = current.Next.Next
			}
		} else {
			current = current.Next
		}

	}
	return head
}
