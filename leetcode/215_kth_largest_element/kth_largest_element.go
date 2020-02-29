// author: kaku
// date: 2019/10/24
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//		leetcode #215 https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
package main

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	pivot := partition(nums, 0, len(nums)-1)

	for pivot+1 != k {
		if pivot < k {
			pivot = partition(nums, pivot+1, len(nums)-1)
		} else {
			pivot = partition(nums, 0, pivot-1)
		}
	}

	return nums[pivot]
}


func partition(arr []int, low, high int) int {
	elem := arr[low]

	for low < high {
		for ; low < high && arr[high] <= elem; high-- {
		}
		arr[low] = arr[high]

		for ; low < high && arr[low] >= elem; low++ {
		}
		arr[high] = arr[low]
	}

	arr[low] = elem
	return low
}

func main() {
	arr := []int{3,2,3,1,2,4,5,5,6}
	fmt.Println(findKthLargest(arr, 4))
}
