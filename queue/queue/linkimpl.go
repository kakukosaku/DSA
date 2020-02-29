// author: kaku
// date: 2020/02/28
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//
package queue

import (
	"errors"
	"fmt"
	"strings"
)

type LinkNode struct {
	Elem interface{}
	Next *LinkNode
}

type LinkQueue struct {
	Front *LinkNode
	Rear  *LinkNode
}

// New init a link queue
func New() *LinkQueue {
	q := LinkQueue{}
	return &q
}

// IsEmpty judge link queue is empty or not
func (q *LinkQueue) IsEmpty() bool {

	if q.Front == nil {
		return true
	}

	return false
}

// EnQueue enqueue an element n
func (q *LinkQueue) EnQueue(n LinkNode) {
	if q.Front == nil {
		q.Front = &n
		q.Rear = &n
	} else {
		q.Rear.Next = &n
		q.Rear = &n
	}

}

// Dequeue dequeue an element n
func (q *LinkQueue) DeQueue() (LinkNode, error) {
	if q.Front == nil {
		return LinkNode{}, errors.New("dequeue failed, queue is empty")
	}

	var n LinkNode = *q.Front
	q.Front = n.Next
	return n, nil
}

func (q *LinkQueue) Show() {
	displayStr := strings.Builder{}
	displayStr.WriteString("Show Queue:\n\t")

	if q.Front == nil {
		displayStr.WriteString("nil\n")
	} else {
		for currNode := q.Front; currNode != nil; currNode = currNode.Next {
			displayStr.WriteString(fmt.Sprintf(" %v ->", currNode.Elem))
		}
		displayStr.WriteString(" end!\n")
	}

	fmt.Println(displayStr.String())
}
