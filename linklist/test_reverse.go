// author: kaku
// date: 2020/03/02
//
// GitHub:
//	https://github.com/kakukosaku
//
// Description:
//
package main

import (
	"github.com/kakukosaku/DSA/linklist/linklist"
	"github.com/kakukosaku/DSA/linklist/reverse"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	l := linklist.New()
	l.AddSlice(s...)
	l.Show()

	reverse.ReverseLinkList(l)
	l.Show()
}
