// author: kaku
// date: 2020/02/28
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//
package stack

import (
	"errors"
	"fmt"
	"strings"
)

type LinkNode struct {
	Elem interface{}
	Next *LinkNode
}

type LinkStack struct {
	top *LinkNode
}

// New init a empty stack
func New() *LinkStack {
	return new(LinkStack)
}

// IsEmpty judge stack is empty or not
func (s *LinkStack) IsEmpty() bool {
	if s.top == nil {
		return true
	}

	return false
}

// PushStack add an elem
func (s *LinkStack) PushStack(n LinkNode) {
	n.Next = s.top
	s.top = &n
}

// PopStack remove an elem
func (s *LinkStack) PopStack() (LinkNode, error) {
	if s.top == nil {
		return LinkNode{}, errors.New("pop stack failed, stack is empty")
	}

	var n LinkNode = *s.top
	s.top = n.Next

	return n, nil
}

func (s *LinkStack) Show() {

	displayStr := strings.Builder{}
	displayStr.WriteString("Show Stack from top->bottom:\n\t")

	if s.top == nil {
		displayStr.WriteString("nil\n")
	} else {
		displayStr.WriteString("<-")
		for currNode := s.top; currNode != nil; currNode = currNode.Next {
			displayStr.WriteString(fmt.Sprintf(" %v|", currNode.Elem))
		}
		displayStr.WriteString("\n")
	}

	fmt.Println(displayStr.String())
}
