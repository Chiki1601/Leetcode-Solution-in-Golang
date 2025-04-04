func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_, res := helper(root)
	return res
}

func helper(node *TreeNode) (int, *TreeNode) {
	if node == nil {
		return 0, nil
	}

	if node.Left == nil && node.Right == nil {
		return 1, node
	}

	leftDepth, leftCommon := helper(node.Left)
	rightDepth, rightCommon := helper(node.Right)
	if leftCommon != nil && rightCommon != nil {
		if leftDepth == rightDepth {
			return rightDepth + 1, node
		} else if leftDepth > rightDepth {
			return leftDepth + 1, leftCommon
		} else {
			return rightDepth + 1, rightCommon
		}
	}

	if leftCommon != nil {
		return leftDepth + 1, leftCommon
	}

	return rightDepth + 1, rightCommon
}
