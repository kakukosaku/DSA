// author: kaku
// date: 2019/02/27
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//
package heap_impl

import "github.com/kakukosaku/DSA/sort/utils"

func heapnify(arr []int, checkIdx, arrSize int) {
	theEle := arr[checkIdx]

	for checkIdx*2+1 < arrSize {
		lChildIdx := checkIdx*2 + 1
		rChildIdx := checkIdx*2 + 2
		biggerIdx := lChildIdx

		if rChildIdx < arrSize && arr[lChildIdx] < arr[rChildIdx] {
			biggerIdx = rChildIdx
		}

		if theEle > arr[biggerIdx] {
			break
		} else {
			arr[checkIdx] = arr[biggerIdx]
			checkIdx = biggerIdx
		}
	}

	arr[checkIdx] = theEle
}

func buildHead(arr []int, arrSize int) {
	for checkIdx := len(arr)/2 - 1; checkIdx >= 0; checkIdx-- {
		heapnify(arr, checkIdx, arrSize)
	}
}

func HeapSort(arr []int) []int {
	buildHead(arr, len(arr))

	for sortedNum := 1; sortedNum < len(arr); sortedNum++ {
		utils.Swap(arr, 0, len(arr)-sortedNum)
		heapnify(arr, 0, len(arr)-sortedNum)
	}

	return arr
}
