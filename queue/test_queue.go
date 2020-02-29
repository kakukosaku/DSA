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
	"fmt"
	"github.com/kakukosaku/DSA/queue/queue"
	"log"
)

func main() {
	q := queue.New()
	q.Show()

	arr := []int{1, 2, 3, 4, 5}
	for _, v := range arr {
		n := queue.LinkNode{Elem: v}
		q.EnQueue(n)
	}

	q.Show()
	for i := 0; i < len(arr)+1; i++ {
		n, err := q.DeQueue()
		fmt.Printf("Got a node(%v)\n", n.Elem)
		if err != nil {
			log.Fatal(err)
		}
		q.Show()
	}

}
