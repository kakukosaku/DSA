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
	HeadPtr *Node
	Size    int
}

func New() *LinkList {
	l := LinkList{HeadPtr: &Node{}, Size: 0}
	return &l
}

// Add head insert
func (l *LinkList) Add(e interface{}) {
	n := Node{Elem: e}
	n.Next = l.HeadPtr.Next
	l.HeadPtr.Next = &n
}

func (l *LinkList) AddSlice(s []interface{}) {
	for _, v := range s {
		l.Add(v)
	}
}

func (l *LinkList) Show() {
	fmt.Printf("Show link list:\n\t")

	for cn := l.HeadPtr.Next; cn != nil; {
		fmt.Printf("%v -> ", cn.Elem)
		cn = cn.Next
	}
	fmt.Println("nil")
}
