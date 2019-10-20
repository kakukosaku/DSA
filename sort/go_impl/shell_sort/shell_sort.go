/*
author: kaku
date: 18/08/05 Sun
github:	https://github.com/kakuchange
description:
	Shell sort go implement.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ShellSort(s []int) {
	increment := len(s) / 2
	for increment > 0 {
		for i := increment; i < len(s); i++ {
			j := i
			for j >= increment && s[j] < s[j-increment] {
				s[j], s[j-increment] = s[j-increment], s[j]
				j = j - increment
			}
		}
		increment = increment / 2
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
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	s := make([]int, 10, 10) // len 10, cap 10 later is optional.
	for index, _ := range s {
		s[index] = rand.Intn(100)
	}
	fmt.Println("init slice:", s)
	ShellSort(s)
	sortedSlice := s
	fmt.Println("sort slice:", sortedSlice)

}
