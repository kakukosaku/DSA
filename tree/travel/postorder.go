// author: kaku
// date: 2020/03/01
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//
package travel

import (
	"container/list"
	"github.com/kakukosaku/DSA/tree"
)

func PostOrder(root *tree.Node) []int {
	travelRest := make([]int, 0)
	if root == nil {
		return travelRest
	}

	stack1 := list.New()
	stack2 := list.New()
	stack1.PushFront(root)
	for stack1.Len() != 0 {
		elem := stack1.Front()
		stack1.Remove(elem)
		node := elem.Value.(*tree.Node)
		stack2.PushFront(node)

		if node.Left != nil {
			stack1.PushFront(node.Left)
		}
		if node.Right != nil {
			stack1.PushFront(node.Right)
		}
	}

	for stack2.Len() != 0 {
		elem := stack2.Front()
		stack2.Remove(elem)
		node := elem.Value.(*tree.Node)
		travelRest = append(travelRest, node.Val)
	}

	return travelRest
}
