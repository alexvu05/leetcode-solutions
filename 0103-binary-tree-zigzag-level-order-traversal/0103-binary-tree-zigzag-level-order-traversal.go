/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    result := [][]int{}
    if root == nil {
        return result
    }
    queue := []*TreeNode{root}
    leftToRight := true  // level 0 goes left→right
    for len(queue) > 0 {
        layerSize := len(queue)
        currentLevel := make([]int, layerSize)
        for i := 0; i < layerSize; i++ {
            node := queue[0]
            queue = queue[1:]
            // Place value at correct position based on direction
            if leftToRight {
                currentLevel[i] = node.Val
            } else {
                currentLevel[layerSize-1-i] = node.Val  // reverse index
            }
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        result = append(result, currentLevel)
        leftToRight = !leftToRight  // flip direction for next level
    }

    return result
}