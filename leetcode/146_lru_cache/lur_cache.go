// author: kaku
// date: 2020/02/29
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//		leetcode #146146146 https://leetcode-cn.com/problems/lru-cache/
package main

import (
	"fmt"
)

type LinkNode struct {
	Next  *LinkNode
	Prior *LinkNode
	Key   int
	Value int
}

type DoubleLinkList struct {
	HeadPtr *LinkNode
	TailPtr *LinkNode
	Size    int
}

func (d *DoubleLinkList) AddFirst(n *LinkNode) {
	n.Next = d.HeadPtr.Next
	n.Prior = d.HeadPtr

	d.HeadPtr.Next.Prior = n
	d.HeadPtr.Next = n
	d.Size++
}

func (d *DoubleLinkList) Remove(n *LinkNode) *LinkNode {
	n.Prior.Next = n.Next
	n.Next.Prior = n.Prior
	d.Size--

	return n
}

func (d *DoubleLinkList) RemoveLast() *LinkNode {
	// 除头尾结点外, 有其它结点
	if d.TailPtr.Prior != d.HeadPtr {
		return d.Remove(d.TailPtr.Prior)
	}

	return nil
}

type LRUCache struct {
	Map      map[int]*LinkNode
	DLL      DoubleLinkList
	Capacity int
}

func (this *LRUCache) Get(key int) int {
	if this.Capacity <= 0 {
		return -1
	}

	if v, ok := this.Map[key]; ok {
		this.Put(key, v.Value)
		return v.Value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if n, ok := this.Map[key]; ok {
		this.DLL.Remove(n)
	} else {
		if this.Capacity == this.DLL.Size {
			n = this.DLL.RemoveLast()
			delete(this.Map, n.Key)
		}
	}

	n := &LinkNode{Key: key, Value:value}
	this.DLL.AddFirst(n)
	this.Map[key] = n
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{Capacity: capacity}
	lru.Map = make(map[int]*LinkNode)
	lru.DLL.HeadPtr = new(LinkNode)
	lru.DLL.TailPtr = new(LinkNode)

	lru.DLL.HeadPtr.Next = lru.DLL.TailPtr
	lru.DLL.TailPtr.Prior = lru.DLL.HeadPtr
	return lru
}

func main() {
	lru := Constructor(1)
	lru.Put(2, 1)
	lru.Put(2, 2)

	fmt.Println(lru.Get(2))
	lru.Put(1, 1)
	lru.Put(4, 1)
	fmt.Println(lru.Get(2))
}
