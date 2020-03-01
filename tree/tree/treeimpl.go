// author: kaku
// date: 2020/02/28
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//
package tree

import "fmt"

type Node struct {
	Elem   int
	LChild *Node
	RChild *Node
}

type Tree struct {
	Root *Node
}

func (n Node) Show() {
	fmt.Printf("\t%v\n", n.Elem)
	if n.LChild != nil && n.RChild != nil {
		fmt.Printf("%v\t\t%v\n", n.LChild.Elem, n.RChild.Elem)
	}

	if n.LChild == nil && n.RChild != nil {
		fmt.Printf("\t\t%v\n", n.RChild.Elem)
	}

	if n.LChild != nil && n.RChild == nil {
		fmt.Printf("%v\n", n.LChild.Elem)
	}

	if n.LChild == nil && n.RChild == nil {
		fmt.Printf("nil\t\tnil\n")
	}
}
