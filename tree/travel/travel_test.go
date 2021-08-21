// author: kaku
// date: 2020/02/28
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//
package travel

import (
	"fmt"
	"github.com/kakukosaku/DSA/tree"
	"testing"
)

func TestTravel(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	root := tree.NewCompleteTree(arr)
	root.Show()

	var travelRest []int
	travelRest = PreOrder(root)
	fmt.Println("Pre order travel:\t", travelRest)
	travelRest = InOrder(root)
	fmt.Println("In order travel:\t", travelRest)
	travelRest = PostOrder(root)
	fmt.Println("Post order travel:\t", travelRest)

	fmt.Println()
	travelRest = PreOrderRecursive(root, make([]int, 0))
	fmt.Println("Pre order travel:\t", travelRest)
	travelRest = InOrderRecursive(root, make([]int, 0))
	fmt.Println("In order travel:\t", travelRest)
	travelRest = PostOrderRecursive(root, make([]int, 0))
	fmt.Println("Post order travel:\t", travelRest)
}
