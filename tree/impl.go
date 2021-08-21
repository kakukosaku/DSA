// author: kaku
// date: 2020/02/28
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//
package tree

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n *Node) Show() {
	queue := list.New()
	if n == nil {
		log.Println("tree is empty")
		return
	}

	buf := bufio.NewWriter(os.Stdout)
	defer func() {
		_ = buf.Flush()
	}()
	_, _ = buf.WriteString("\nLevel show tree:")

	currLevel := 0
	queue.PushBack(n)
	for queue.Len() != 0 {
		_, _ = buf.WriteString(fmt.Sprintf("\nLv.%d:\t", currLevel))
		currLevel++
		for levelNum := queue.Len(); levelNum > 0; levelNum-- {
			elem := queue.Front()
			queue.Remove(elem)
			node := elem.Value.(*Node)
			_, _ = buf.WriteString(fmt.Sprintf("%d ", node.Val))

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	_, _ = buf.WriteString("\n\n")
}
