// Author: kaku
// Date: 2021/8/21
//
// GitHub:
//	https://github.com/kakukosaku
//
package tree

import (
	"testing"
)

func TestNewCompleteTree(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	root := NewCompleteTree(s)

	root.Show()
}
