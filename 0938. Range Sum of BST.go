/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    if root == nil {
        return 0
    }

    res := root.Val
    
    if res < low || res > high {
        res = 0
    }

    return res + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}
