/*
author: kaku
date: 18/08/06 Mon
github:	https://github.com/kakuchange
description:
	This module do what.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func MergeSort(s []int) []int {
	n := len(s)
	if n < 2 {
		return s
	}
	key := n / 2
	left := MergeSort(s[0:key])
	right := MergeSort(s[key:])
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	tmp := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			tmp = append(tmp, left[i])
			i++
		} else {
			tmp = append(tmp, right[j])
			j++
		}
	}
	tmp = append(tmp, left[i:]...)
	tmp = append(tmp, right[j:]...)
	return tmp
}
func main() {
	fmt.Println("vim-go")
	// init a slice
	// var s []int
	// for i := 0; i < 10; i++ {
	// 	s = append(s, rand.Intn(100))
	// }

	// or you and do like this
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	s := make([]int, 10, 10) // len 10, cap 10 later is optional.
	for index, _ := range s {
		s[index] = rand.Intn(100)
	}
	fmt.Println("init slice:", s)
	sortedSlice := MergeSort(s)
	fmt.Println("sort slice:", sortedSlice)

}
