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
	"github.com/kakukosaku/DSA/stack/stack"
	"log"
)

func main() {
	s := stack.New()
	s.Show()

	arr := []int{1, 2, 3, 4, 5}

	for _, v := range arr {
		s.PushStack(stack.LinkNode{Elem: v})
	}

	s.Show()
	for i := 0; i < len(arr)+1; i++ {
		n, err := s.PopStack()
		if err != nil {
			log.Fatal(n)
		}
		s.Show()
	}
}
