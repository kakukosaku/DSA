package travel

import (
	"fmt"
	"github.com/kakukosaku/DSA/tree/tree"
)

func PreOrderRecursive(t *tree.Tree) {
	fmt.Printf("Tree pre-order recursive travel:\n\t")
	preOrderRecursive(t.Root)
	fmt.Println()
}

func preOrderRecursive(n *tree.Node) {
	if n != nil {
		fmt.Printf("%v -> ", n.Elem)
		preOrderRecursive(n.LChild)
		preOrderRecursive(n.RChild)
	}
}

func InOrderRecursive(t *tree.Tree) {
	fmt.Printf("Tree in-order recursive travel:\n\t")
	inOrderRecursive(t.Root)
	fmt.Println()
}

func inOrderRecursive(n *tree.Node) {
	if n != nil {
		inOrderRecursive(n.LChild)
		fmt.Printf("%v -> ", n.Elem)
		inOrderRecursive(n.RChild)
	}
}

func PostOrderRecursive(t *tree.Tree) {
	fmt.Printf("Tree post-order recursive travel:\n\t")
	postOrderRecursive(t.Root)
	fmt.Println()
}

func postOrderRecursive(n *tree.Node) {
	if n != nil {
		postOrderRecursive(n.LChild)
		postOrderRecursive(n.RChild)
		fmt.Printf("%v -> ", n.Elem)
	}
}