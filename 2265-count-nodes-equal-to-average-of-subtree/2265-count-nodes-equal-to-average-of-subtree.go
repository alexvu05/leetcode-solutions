/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfSubtree(root *TreeNode) int {
    count := 0
    var dfs func(node *TreeNode) (int, int) // returns (sum, nodeCount)
    dfs = func(node *TreeNode) (int, int) {
        if node == nil {
            return 0, 0
        }
        leftSum, leftCnt   := dfs(node.Left)
        rightSum, rightCnt := dfs(node.Right)
        totalSum := leftSum + rightSum + node.Val
        totalCnt := leftCnt + rightCnt + 1
        // Check if node value equals floor average of its subtree
        if node.Val == totalSum/totalCnt {
            count++
        }
        return totalSum, totalCnt
    }
    dfs(root)
    return count
}