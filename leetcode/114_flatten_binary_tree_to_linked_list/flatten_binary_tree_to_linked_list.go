package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)
	if root.Left != nil {
		lastNode := root.Left
		for lastNode.Right != nil {
			lastNode = lastNode.Right
		}
		lastNode.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}

}

func flattenNonRecursive(root *TreeNode) {
	curr := root
	for curr != nil {
		if root.Left != nil {
			nextNode := root.Left
			prevNode := nextNode
			for prevNode.Right != nil {
				prevNode = nextNode.Right
			}
			prevNode.Right = curr.Right
			curr.Right = nextNode
			curr.Left = nil
		}
		curr = curr.Right
	}
}
