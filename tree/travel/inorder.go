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

func InOrder(root *tree.Node) []int {
	travelRest := make([]int, 0)
	if root == nil {
		return travelRest
	}

	stack := list.New()
	for currNode := root; currNode != nil; {
		stack.PushFront(currNode)
		currNode = currNode.Left
	}
	for stack.Len() != 0 {
		elem := stack.Front()
		stack.Remove(elem)
		node := elem.Value.(*tree.Node)
		travelRest = append(travelRest, node.Val)

		if node.Right != nil {
			for currNode := node.Right; currNode != nil; {
				stack.PushFront(currNode)
				currNode = currNode.Left
			}
		}
	}

	return travelRest
}
