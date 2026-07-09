/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }
    // Step 1: find the middle node using fast & slow pointers
    middle := findMiddle(head)
    // Step 2: reverse the second half of the list
    secondHalf := reverse(middle.Next)
    // Step 3: compare first half and reversed second half
    p1 := head
    p2 := secondHalf
    result := true
    for p2 != nil {
        if p1.Val != p2.Val {
            result = false
            break
        }
        p1 = p1.Next
        p2 = p2.Next
    }
    // Restore the list (optional but good practice in real systems)
    middle.Next = reverse(secondHalf)
    return result
}

// findMiddle returns the middle node (left-middle for even length)
func findMiddle(head *ListNode) *ListNode {
    slow := head
    fast := head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}

// reverse reverses a linked list and returns the new head
func reverse(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}