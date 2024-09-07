package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	} else if compare(elem, root.Data) > 0 {
		return BTreeSearchItem(root.Right, elem)
	} else if compare(elem, root.Data) < 0 {
		return BTreeSearchItem(root.Left, elem)
	} else {
		return root
	}
}
