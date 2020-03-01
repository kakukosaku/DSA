// author: kaku
// date: 2020/02/28
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//
package main

import (
	"fmt"
	"github.com/kakukosaku/DSA/tree/maxdepth"
	"github.com/kakukosaku/DSA/tree/utils"
)

func main() {
	var arrSize = 7
	var arr []int

	for i := 0; i < arrSize; i++ {
		arr = append(arr, i)
	}

	t := utils.NewCompleteTreeFromSlice(arr)
	fmt.Printf("Max depth calculate used recursive DFS:\n\t%v\n", maxdepth.DFSMaxDepthRecursive(t.Root))
	fmt.Printf("Max depth calculate used queue base BFS:\n\t%v\n", maxdepth.BFSMaxDepth(t.Root))
}
