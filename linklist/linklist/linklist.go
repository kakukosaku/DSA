// author: kaku
// date: 2020/03/01
//
// GitHub:
//	https://github.com/kakukosaku
//
// Description:
//	单链表相关操作
package linklist

import "fmt"

type Node struct {
	Elem interface{}
	Next *Node
}

type LinkList struct {
	Head *Node
	Tail *Node
	Size int
}

func New() *LinkList {
	head := &Node{}
	tail := &Node{}

	head.Next = tail
	l := LinkList{
		Head: head,
		Tail: tail,
		Size: 0,
	}

	return &l
}

func (l *LinkList) Empty() bool {
	if l.Head.Next == l.Tail {
		return true
	}
	return false
}

// Add head insert
func (l *LinkList) Add(e interface{}) {
	n := Node{Elem: e}
	n.Next = l.Head.Next
	l.Head.Next = &n
	l.Size++
}

func (l *LinkList) AddSlice(s ...int) {
	for _, v := range s {
		l.Add(v)
	}
}

func (l *LinkList) Show() {
	fmt.Printf("Show link list:\n\tHeadNode -> ")

	for cn := l.Head.Next; cn.Next != nil; {
		fmt.Printf("%v -> ", cn.Elem)
		cn = cn.Next
	}

	fmt.Println("TailNode!")
}
