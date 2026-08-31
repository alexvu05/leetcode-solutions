/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nodesBetweenCriticalPoints(head *ListNode) []int {
    if head == nil || head.Next == nil || head.Next.Next == nil {
        return []int{-1, -1}
    }
    minDist  := int(^uint(0) >> 1)  // math.MaxInt
    firstIdx := -1
    lastIdx  := -1
    prevIdx  := -1
    prev := head
    curr := head.Next
    idx  := 1  // curr is at index 1 (0-indexed)
    for curr.Next != nil {
        next := curr.Next
        // Check if curr is a critical point
        isCritical := (curr.Val > prev.Val && curr.Val > next.Val) ||
                      (curr.Val < prev.Val && curr.Val < next.Val)
        if isCritical {
            if firstIdx == -1 {
                firstIdx = idx
            } else {
                // Update minDist with distance from previous critical point
                if idx-prevIdx < minDist {
                    minDist = idx - prevIdx
                }
            }
            prevIdx = idx
            lastIdx = idx
        }
        prev = curr
        curr = next
        idx++
    }
    // Need at least 2 critical points
    if firstIdx == lastIdx {
        return []int{-1, -1}
    }
    return []int{minDist, lastIdx - firstIdx}
}