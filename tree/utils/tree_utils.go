// author: kaku
// date: 2020/02/28
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//
package utils

import "github.com/kakukosaku/DSA/tree/tree"

func New() *tree.Tree {
	return new(tree.Tree)
}

func NewCompleteTreeFromSlice(s []int) *tree.Tree {
	t := New()
	if len(s) == 0 {
		return t
	}

	waitingToFullNode := make([]*tree.Node, 0)
	rootElem, s := PopFromIntSlice(s, 0)
	n := tree.Node{Elem: rootElem}
	t.Root = &n
	waitingToFullNode = append(waitingToFullNode, t.Root)

	for _, v := range s {
		checkNode := waitingToFullNode[0]

		newNode := tree.Node{Elem: v}
		if checkNode.LChild == nil {
			checkNode.LChild = &newNode
		} else {
			checkNode.RChild = &newNode
			waitingToFullNode = waitingToFullNode[1:]
		}

		waitingToFullNode = append(waitingToFullNode, &newNode)
	}

	return t
}

func PopFromIntSlice(s []int, idx int) (int, []int) {
	ele := s[idx]
	changedSlice := append(s[:idx], s[idx+1:]...)
	return ele, changedSlice
}
