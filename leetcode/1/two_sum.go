// author: kaku
// date: 2019/10/24
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//		leetcode #1 https://leetcode.com/problems/two-sum/

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		} else {
			m[v] = i
		}
	}
	return []int{0, 0}
}

func main() {
	s := []int{2, 7, 11, 15}
	y := twoSum(s, 18)
	fmt.Println(y)
}
