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
	arrInterface := make([]interface{}, len(arr))
	for i, v := range arr {
		arrInterface[i] = v
	}

	l := linklist.New()
	l.AddSlice(arrInterface)

	l.Show()
}
