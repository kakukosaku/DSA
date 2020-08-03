/*
author: kaku
date: 18/08/14 Tue
github:	https://github.com/kakuchange
description:
	This module do what.
*/

package reverse

import (
	"github.com/kakukosaku/DSA/linklist/linklist"
)

func ReverseLinkList(l *linklist.LinkList) {
	prevNodePtr := &linklist.Node{}
	currNodePtr := l.Head
	l.Tail = l.Head
	for currNodePtr != nil {
		nextNodePtr := currNodePtr.Next
		currNodePtr.Next = prevNodePtr
		prevNodePtr = currNodePtr
		currNodePtr = nextNodePtr
	}
	l.Head = prevNodePtr
	// Important!
	l.Tail.Next = nil
}
