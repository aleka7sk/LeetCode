package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}
	i := 1
	current := head
	prev := &ListNode{}
	next := &ListNode{}
	head = current.Next
	for current != nil {
		if i%2 == 0 {
			next = current.Next
			current.Next = prev
			prev.Next = next
			current = next
		} else {
			if current.Next == nil {
				prev.Next = current
				break
			}
			prev.Next = current.Next
			prev = current
			current = current.Next
		}
		i++
	}
	return head
}
