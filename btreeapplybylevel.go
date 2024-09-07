package piscine

// look for a better approach

type queue []string

func treeRoot(node *TreeNode) *TreeNode {
	curr := node
	for ; curr.Parent != nil; curr = curr.Parent {
	}
	return curr
}

func applyByLevel(node *TreeNode, nodesData queue, f func(...interface{}) (int, error)) {
	if node == nil {
		return
	}
	f(node.Data)
	nodesData = nodesData[1:]
	if node.Left != nil {
		nodesData = append(nodesData, node.Left.Data)
	}
	if node.Right != nil {
		nodesData = append(nodesData, node.Right.Data)
	}
	// There's no nodes left.
	if len(nodesData) == 0 {
		return
	}
	root := treeRoot(node)
	nextNode := BTreeSearchItem(root, nodesData[0])
	applyByLevel(nextNode, nodesData, f)
}

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		var nodesToFindAndApply queue = []string{root.Data}
		applyByLevel(root, nodesToFindAndApply, f)
	}
}
