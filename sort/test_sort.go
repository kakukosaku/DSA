// author: kaku
// date: 2019/02/27
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//
package main

import (
	"github.com/kakukosaku/DSA/sort/bubble"
	"github.com/kakukosaku/DSA/sort/insert"
	"github.com/kakukosaku/DSA/sort/quick"
	"github.com/kakukosaku/DSA/sort/utils"
	selectimpl "github.com/kakukosaku/DSA/sort/select_impl"
	heapimpl "github.com/kakukosaku/DSA/sort/heap_impl"
)

var ArraySize int = 10

func testCase(arr []int, arrSize int, description string, sortFunc func([]int) []int) {
	_arr := make([]int, arrSize)
	copy(_arr, arr)
	_arr = sortFunc(_arr)
	utils.Show(_arr, description)
	utils.IsSorted(_arr, false)
}

func main() {
	arr := utils.RandomArray(ArraySize)
	utils.Show(arr, "Original Array")

	testCase(arr, len(arr), "After bubble sort", bubble.BubbleSort)
	testCase(arr, len(arr), "After insert sort", insert.InsertSort)
	testCase(arr, len(arr), "After select sort", selectimpl.SelectSort)
	testCase(arr, len(arr), "After heap sort", heapimpl.HeapSort)
	testCase(arr, len(arr), "After quick sort", quick.QuickSortTestWrap)
}
