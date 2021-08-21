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

func PreOrder(root *tree.Node) []int {
	travelRest := make([]int, 0)
	if root == nil {
		return travelRest
	}

	stack := list.New()
	stack.PushFront(root)
	for stack.Len() != 0 {
		elem := stack.Front()
		node := elem.Value.(*tree.Node)
		stack.Remove(elem)
		travelRest = append(travelRest, node.Val)

		if node.Right != nil {
			stack.PushFront(node.Right)
		}
		if node.Left != nil {
			stack.PushFront(node.Left)
		}
	}

	return travelRest
}
