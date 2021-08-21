// author: kaku
// date: 2020/02/28
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//
package tree

import "container/list"

func New(rootVal int) *Node {
	root := new(Node)
	root.Val = rootVal
	return root
}

func NewCompleteTree(s []int) *Node {
	if len(s) == 0 {
		return nil
	}

	root := New(s[0])
	notFullNodeQueue := list.New()
	notFullNodeQueue.PushBack(root)

	for _, v := range s[1:] {
		node := notFullNodeQueue.Front().Value.(*Node)
		newNode := &Node{Val: v}

		if node.Left == nil {
			node.Left = newNode
		} else {
			node.Right = newNode
			notFullNodeQueue.Remove(notFullNodeQueue.Front())
		}
		notFullNodeQueue.PushBack(newNode)
	}
	return root
}
