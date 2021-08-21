package travel

import (
	"github.com/kakukosaku/DSA/tree"
)

func PreOrderRecursive(node *tree.Node, travelRest []int) []int {
	if node == nil {
		return travelRest
	}
	travelRest = append(travelRest, node.Val)
	travelRest = PreOrderRecursive(node.Left, travelRest)
	travelRest = PreOrderRecursive(node.Right, travelRest)
	return travelRest
}

func InOrderRecursive(node *tree.Node, travelRest []int) []int {
	if node == nil {
		return travelRest
	}
	travelRest = InOrderRecursive(node.Left, travelRest)
	travelRest = append(travelRest, node.Val)
	travelRest = InOrderRecursive(node.Right, travelRest)
	return travelRest
}

func PostOrderRecursive(node *tree.Node, travelRest []int) []int {
	if node == nil {
		return travelRest
	}
	travelRest = PostOrderRecursive(node.Left, travelRest)
	travelRest = PostOrderRecursive(node.Right, travelRest)
	travelRest = append(travelRest, node.Val)
	return travelRest
}
