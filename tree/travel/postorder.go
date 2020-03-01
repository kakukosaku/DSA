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

func PostOrder(t *tree.Tree) {
	if t.Root == nil {
		log.Fatal("tree is empty")
	}

	s1 := stack.New()
	s2 := stack.New()
	s1.PushStack(stack.LinkNode{Elem: t.Root})
	for !s1.IsEmpty() {
		sn, err := s1.PopStack()
		if err != nil {
			log.Fatal("stack is empty already")
		}
		tn := sn.Elem.(*tree.Node)

		s2.PushStack(stack.LinkNode{Elem: tn})
		if tn.LChild != nil {
			s1.PushStack(stack.LinkNode{Elem: tn.LChild})
		}
		if tn.RChild != nil {
			s1.PushStack(stack.LinkNode{Elem: tn.RChild})
		}
	}
	fmt.Printf("Tree post-order travel\n\t")
	for !s2.IsEmpty() {
		sn, err := s2.PopStack()
		if err != nil {
			log.Fatal("stack is empty already")
		}
		tn := sn.Elem.(*tree.Node)
		fmt.Printf("%v -> ", tn.Elem)
	}
	fmt.Println()
}
