/*
auther: kaku
date: 18/08/05 Sun
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

func QuickSort(s []int) {
	if len(s) <= 1 {
		return
	}
	//每次选最左边为pivot.
	mid := s[0]
	head, tail := 0, len(s)-1
	for i := 1; i <= tail; {
		if s[i] > mid {
			s[i], s[tail] = s[tail], s[i]
			tail--
		} else {
			s[i], s[head] = s[head], s[i]
			head++
			i++
		}
	}
	quickSort(s[:head])
	quickSort(s[head+1:])
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
	QuickSort(s)
	sortedSlice := s
	fmt.Println("sort slice:", sortedSlice)
}
