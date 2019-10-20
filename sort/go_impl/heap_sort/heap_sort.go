/*
author: kaku
date: 18/08/04 Sat
github:	https://github.com/kakuchange
description:
	Heap sort go implement.
*/

package main

import (
	"fmt"
	"math/rand"
)

func heapnify(s []int, start, end int) {
	leftChild := 2*start + 1
	if leftChild > end {
		return
	}
	smallerChild := leftChild
	rightChild := 2*start + 2
	// 把最大的取出, 逐渐置于堆顶, 取其至数组底部, 实现从小到大
	if rightChild <= end && s[leftChild] < s[rightChild] {
		smallerChild = rightChild
	}
	s[start], s[smallerChild] = s[smallerChild], s[start]
	heapnify(s, smallerChild, end)
}

func HeapSort(s []int) {
	for start, end := len(s)/2, len(s)-1; start > 0; start-- {
		heapnify(s, start, end)
	}
	for i := len(s) - 1; i > 0; i-- {
		s[0], s[i] = s[i], s[0]
		heapnify(s, 0, i-1)
	}
}

func main() {
	fmt.Println("vim-go")
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
	HeapSort(s)
	sortedSlice := s
	fmt.Println("sort slice:", sortedSlice)
}
