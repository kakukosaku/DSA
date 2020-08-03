// author: kaku
// date: 2020/月/日
//
// GitHub:
//	https://github.com/kakukosaku
//
// Description:
//
package main

import (
	"github.com/kakukosaku/DSA/linklist/linklist"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	l := linklist.New()
	l.AddSlice(arr...)
	l.Show()

	l2 := linklist.New()
	l2.Show()
}
