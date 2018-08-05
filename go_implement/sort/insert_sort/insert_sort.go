/*
auther: kaku
date: 年/月/日 Sun
github:	https://github.com/kakuchange
description:
	insert sort
*/

package main

import (
	"fmt"
	"math/rand"
)

func InsertSort(s []int) []int {
	for i := 0; i < len(s); i++ {
		for j := i; j > 0; j-- {
			// 从小到大
			if s[j] < s[j-1] {
				s[j], s[j-1] = s[j-1], s[j]
			} else {
				// 因前面元素已有序, 找到第一个不满足即可退出
				break
			}
		}
	}
	return s
}

func main() {
	// init a slice
	// var s []int
	// for i := 0; i < 10; i++ {
	// 	s = append(s, rand.Intn(100))
	// }

	// or you and do like this
	s := make([]int, 10, 10) // len 10, cap 10 later is optional.
	for index, _ := range s {
		s[index] = rand.Intn(100)
	}
	fmt.Println("init slice:", s)
	sortedSlice := InsertSort(s)
	fmt.Println("sort slice:", sortedSlice)

}
