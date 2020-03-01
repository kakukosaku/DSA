// author: kaku
// date: 2020/02/29
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//		leetcode #704 https://leetcode-cn.com/problems/binary-search/
package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if target < nums[0] || target > nums[len(nums)-1] {
		return -1
	}

	low := 0
	high := len(nums) - 1
	for low <= high {
		middle := low + (high-low)/2
		if nums[middle] > target {
			high = middle - 1
		} else if nums[middle] < target {
			low = middle + 1
		} else {
			return middle
		}
	}

	return -1
}

func main() {
	arr := []int{0, 1, 2, 3, 4}

	fmt.Println(search(arr, 4))
}
