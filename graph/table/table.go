// author: kaku
// date: 2020/月/日
//
// GitHub:
//	https://github.com/kakukosaku
//
// Description:
//
//	图的邻接表法
package table

// Acme 用一维`slice`存储顶点信息
type Acme []acme

// acme 每个顶点都有一个指向边链表的指针
type acme struct {
	Data   int
	ArcPtr *Arc
}

// Arc 用链表存储每个顶点的边信息
type Arc struct {
	Data int
	Next *Arc
}
