package leetcode

/**
 * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	current := list1
	for current.Next != nil {
		current = current.Next
	}
	current.Next = list2
	current_two := list1
	for i := current_two; i != nil; i = i.Next {
		for j := current_two; j != nil; j = j.Next {
			if i.Val < j.Val {
				i.Val, j.Val = j.Val, i.Val
			}
		}
	}
	return list1
}
