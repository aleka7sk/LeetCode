package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    if head == nil {
        return nil
    } 
    count := 0
    cur := head
    for cur != nil {
        cur = cur.Next
        count++
    }
    mid := count / 2
    cur2 := head
    count2 := 0
    for cur2 != nil {
        if count2 == mid {
            return cur2
        }
        cur2 = cur2.Next  
        count2++
    }
    return nil
}
