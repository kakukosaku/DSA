/*
auther: kaku
date: 18/08/04 Sat
github:	https://github.com/kakuchange
description:
	This module do what.
*/

package main

import (
	"fmt"
	"math/rand"
)

func SelectSort(s []int) []int {
	for i := 0; i < len(s); i++ {
		smallIndex := i
		for j := i; j < len(s)-1; j++ {
			// 从小到大
			if s[j+1] < s[j] {
				smallIndex = j + 1
			} else {
				continue
			}
		}
		s[i], s[smallIndex] = s[smallIndex], s[i]
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
	sortedSlice := SelectSort(s)
	fmt.Println("sort slice:", sortedSlice)

}
