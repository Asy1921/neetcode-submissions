/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
    res := []int{}
	var postOrder func(node *TreeNode)
	postOrder = func(node *TreeNode){
		if node==nil{
			return
		}
		// Process Left First
		postOrder(node.Left)
		// Then Right
		postOrder(node.Right)
		// Then Root
		res=append(res,node.Val)
	}
	postOrder(root)
	return res
}
