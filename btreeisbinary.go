package piscine

// look for a better approach
func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	} else if (root.Left == nil || compare(root.Data, root.Left.Data) > 0) &&
		(root.Right == nil || compare(root.Data, root.Right.Data) < 0) {
		return BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
	} else {
		return false
	}
}
