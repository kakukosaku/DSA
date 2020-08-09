// author: kaku
// date: 2019/02/29
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//
package quick

func QuickSortTestWrap(arr []int) []int {
	return QuickSort(arr, 0, len(arr)-1)
}

// QuickSort quick sort use low base pivot
func QuickSort(arr []int, low, high int) []int {
	if low < high {
		pivot := partition(arr, low, high)
		QuickSort(arr, low, pivot - 1)
		QuickSort(arr, pivot + 1, high)
	}

	return arr
}

func partition(arr []int, low, high int) int {
	elem := arr[low]

	for low < high {
		for ; low < high && arr[high] >= elem; high-- {
		}
		arr[low] = arr[high]

		for ; low < high && arr[low] <= elem; low++ {
		}
		arr[high] = arr[low]
	}

	arr[low] = elem
	return low
}
