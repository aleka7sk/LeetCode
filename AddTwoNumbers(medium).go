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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var c1 *ListNode
	var c2 *ListNode
	for c1, c2 = l1, l2; c1 != nil && c2 != nil; c1, c2 = c1.Next, c2.Next {
		if c1.Next != nil && c2.Next == nil {
			c2.Next = &ListNode{Val: 0}
		} else if c2.Next != nil && c1.Next == nil {
			c1.Next = &ListNode{Val: 0}
		}
		if c1 != nil {
			c1.Val = c1.Val + c2.Val
		}
		if c1.Next != nil {
			if c1.Val >= 10 {
				c1.Val = c1.Val - 10
				c1.Next.Val += 1
			}
		} else {
			if c1.Val >= 10 {
				c1.Val = c1.Val - 10
				c1.Next = &ListNode{Val: 1}
			}
		}
	}
	return l1
}
