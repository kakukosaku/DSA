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

func InOrder(t *tree.Tree) {
	if t.Root == nil {
		log.Fatal("tree is empty")
	}

	s := stack.New()
	for cn := t.Root; cn != nil; {
		s.PushStack(stack.LinkNode{Elem: cn})
		cn = cn.LChild
	}
	fmt.Printf("Tree in-order travel:\n\t")
	for !s.IsEmpty() {
		qn, err := s.PopStack()
		if err != nil {
			log.Fatal("stack is empty already")
		}

		tn := qn.Elem.(*tree.Node)
		fmt.Printf("%v -> ", tn.Elem)
		if tn.RChild != nil {
			for cn := tn.RChild; cn != nil; {
				s.PushStack(stack.LinkNode{Elem: cn})
				cn = cn.LChild
			}
		}
	}
	fmt.Println()
}
