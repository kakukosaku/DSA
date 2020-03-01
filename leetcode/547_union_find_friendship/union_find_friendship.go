// author: kaku
// date: 2020/02/29
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//		leetcode #547 https://leetcode-cn.com/problems/friend-circles/
package main

import "fmt"

type UnionSet struct {
	Set []int
}

func (u *UnionSet) Initial(num int) {
	if u.Set == nil {
		u.Set = make([]int, num)
		for i := range u.Set {
			u.Set[i] = i
		}
	}
}

func (u *UnionSet) Find(e int) int {
	if e == u.Set[e] {
		return u.Set[e]
	} else {
		// 路径优化, 实质所有子集合元素均挂在 "子集合的根元素" 下
		u.Set[e] = u.Find(u.Set[e])
		return u.Set[e]
	}
}

func (u *UnionSet) Union(e1, e2 int) {
	p1Id := u.Find(e1)
	p2Id := u.Find(e2)
	u.Set[p2Id] = p1Id

	// refresh
	//u.Find(e2)
}

func findCircleNum(M [][]int) int {
	n := len(M[0])
	u := new(UnionSet)
	u.Initial(n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if M[i][j] != 0 {
				u.Union(i, j)
			}
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		if u.Set[i] == i {
			count++
		}
	}

	return count
}

func main() {
	Fs := [][]int{
		{1, 1, 0},
		{1, 1, 1},
		{0, 0, 1},
	}

	fmt.Println(findCircleNum(Fs))
}
