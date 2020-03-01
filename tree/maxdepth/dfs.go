// author: kaku
// date: 2020/02/28
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//
package maxdepth

import (
	"github.com/kakukosaku/DSA/tree/tree"
)

// NodeMaxDepth recursive dfs get max depth
func DFSMaxDepthRecursive(t *tree.Node) int {
	if t == nil {
		return 0
	}

	lMaxDepth := DFSMaxDepthRecursive(t.LChild)
	rMaxDepth := DFSMaxDepthRecursive(t.RChild)

	maxNum := lMaxDepth
	if rMaxDepth > lMaxDepth {
		maxNum = rMaxDepth
	}

	return maxNum + 1
}


// DFSMaxDepth stack base dfs
func DFSMaxDepth(t *tree.Node) int {
	panic("not implement error")
}


