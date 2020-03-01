// author: kaku
// date: 2020/03/01
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//
package travel

import (
	"fmt"
	"github.com/kakukosaku/DSA/stack/stack"
	"github.com/kakukosaku/DSA/tree/tree"
	"log"
)

func PreOrder(t *tree.Tree) {
	if t.Root == nil {
		log.Fatal("tree is empty")
	}

	currNode := t.Root
	s := stack.New()
	s.PushStack(stack.LinkNode{Elem: currNode})
	fmt.Printf("Tree pre-order travel:\n\t")
	for !s.IsEmpty() {
		qn, err := s.PopStack()
		if err != nil {
			log.Fatal("stack is empty already")
		}

		tn := qn.Elem.(*tree.Node)
		fmt.Printf("%v -> ", tn.Elem)

		if tn.RChild != nil {
			s.PushStack(stack.LinkNode{Elem: tn.RChild})
		}

		if tn.LChild != nil {
			s.PushStack(stack.LinkNode{Elem: tn.LChild})
		}
	}
	fmt.Println()
}
