/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
    ans := root.Val
    queue := []*TreeNode{root}
    for len(queue) > 0{
        temp := []*TreeNode{}
        for _, node := range queue{
            if node.Left != nil{
                temp = append(temp, node.Left)
            }
            if node.Right != nil{
                temp = append(temp, node.Right)
            }
        }
        if len(temp) > 0{
            ans = temp[0].Val
        }
        queue = temp
    }
    return ans
}
