// author: kaku
// date: 2020/02/28
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//
package main

import (
	"github.com/kakukosaku/DSA/tree/utils"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	t := utils.NewCompleteTreeFromSlice(arr)

	t.Root.Show()
	t.Root.LChild.Show()
	t.Root.RChild.Show()
}
