package travel

import (
	"fmt"
	"github.com/kakukosaku/DSA/tree/tree"
)

func PreorderRecursive(t *tree.Tree) {
	fmt.Printf("Tree pre-order recursive travel:\n\t")
	preorderRecursive(t.Root)
	fmt.Println()
}

func preorderRecursive(n *tree.Node) {
	if n != nil {
		fmt.Printf("%v -> ", n.Elem)
		preorderRecursive(n.LChild)
		preorderRecursive(n.RChild)
	}
}