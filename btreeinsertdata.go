package piscine

import "fmt"

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	for curr := root; curr != nil; {
		fmt.Println("data: ", curr.Data)
		if Compare(data, curr.Data) > 0 {
		// if Compare(data, root.Data) > 0 {
			if curr.Right != nil {
				curr = curr.Right
			} else {
				curr.Right = &TreeNode{curr, nil, nil, data}
				return curr
			}
		} else {
			if curr.Left != nil {
				curr = curr.Left
			} else {
				curr.Left = &TreeNode{curr, nil, nil, data}
				return curr
			}
		}
	}

	root = &TreeNode{nil, nil, nil, data}
	return root
}
