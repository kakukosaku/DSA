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
	var prev *linklist.Node
	curr := l.HeadPtr.Next
	for curr != nil {
		tmpNode := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmpNode
	}

	l.HeadPtr.Next = prev
}
