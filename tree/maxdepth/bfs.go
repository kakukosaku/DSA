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
	"github.com/kakukosaku/DSA/queue/queue"
	"github.com/kakukosaku/DSA/tree/tree"
)

// BFSMaxDepth queue base bfs
func BFSMaxDepth(t *tree.Node) int {
	type E struct {
		Node  *tree.Node
		Depth int
	}

	q := queue.New()
	maxDepth := 0
	if t != nil {
		q.EnQueue(queue.LinkNode{Elem: E{t.LChild, maxDepth + 1}})
		q.EnQueue(queue.LinkNode{Elem: E{t.RChild, maxDepth + 1}})
	}

	for !q.IsEmpty() {
		n, err := q.DeQueue()
		if err != nil {
			break
		}

		e := n.Elem.(E)
		if e.Depth > maxDepth {
			maxDepth = e.Depth
		}

		if e.Node != nil {
			q.EnQueue(queue.LinkNode{Elem: E{e.Node.LChild, e.Depth + 1}})
			q.EnQueue(queue.LinkNode{Elem: E{e.Node.RChild, e.Depth + 1}})
		}
	}

	return maxDepth
}
