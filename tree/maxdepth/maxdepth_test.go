// author: kaku
// date: 2020/02/28
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//
package maxdepth_test

import (
	"fmt"
	"github.com/kakukosaku/DSA/tree"
	"github.com/kakukosaku/DSA/tree/maxdepth"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	var arrSize = 7
	var arr []int

	for i := 0; i < arrSize; i++ {
		arr = append(arr, i)
	}

	root := tree.NewCompleteTree(arr)
	fmt.Printf("Max depth calculate used recursive DFS:\n\t%v\n", maxdepth.DFSMaxDepth(root))
	fmt.Printf("Max depth calculate used queue base BFS:\n\t%v\n", maxdepth.BFSMaxDepth(root))
}
