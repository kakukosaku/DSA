// author: kaku
// date: 2020/02/29
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//		leetcode #384 https://leetcode-cn.com/problems/shuffle-an-array/
package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	original []int
	shuffled []int
}

func Constructor(nums []int) Solution {
	shuffled := make([]int, len(nums))
	copy(shuffled, nums)
	return Solution{
		original: nums,
		shuffled: shuffled,
	}
}

/** Resets the array to its original configuration and return it. */
func (s *Solution) Reset() []int {
	return s.original
}

/** Returns a random shuffling of the array. */
func (s *Solution) Shuffle() []int {
	for i := range s.shuffled {
		idx := i + rand.Intn(len(s.shuffled) - i)
		s.shuffled[i], s.shuffled[idx] = s.shuffled[idx], s.shuffled[i]
	}
	return s.shuffled
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	obj := Constructor(nums)
	param_1 := obj.Reset();
	param_2 := obj.Shuffle();
	fmt.Printf("%#v\n", param_1)
	fmt.Printf("%#v\n", param_2)
	param_1 = obj.Reset();
	param_2 = obj.Shuffle();
	fmt.Printf("%#v\n", param_1)
	fmt.Printf("%#v\n", param_2)
}
