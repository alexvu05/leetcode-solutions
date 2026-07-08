/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    pA := headA
    pB := headB
    // Each pointer traverses at most m+n nodes before meeting
    // At intersection: pA == pB (same node reference)
    // No intersection: pA == pB == nil (both reach the end simultaneously)
    for pA != pB {
        if pA == nil {
            pA = headB  // switch to list B after exhausting list A
        } else {
            pA = pA.Next
        }
        if pB == nil {
            pB = headA  // switch to list A after exhausting list B
        } else {
            pB = pB.Next
        }
    }
    return pA  // either the intersection node or nil
}