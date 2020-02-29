// author: kaku
// date: 2020/02/27
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//
package bubble

import "github.com/kakukosaku/DSA/sort/utils"

// BubbleSort impl
func BubbleSort(arr []int) []int {
	var sortedNum int = 0
	var moved bool = true

	for sortedNum < len(arr) {
		if moved {
			for checkIdx := 0; checkIdx < len(arr)-sortedNum-1; checkIdx ++ {
				// small -> big
				if arr[checkIdx] > arr[checkIdx+1] {
					utils.Swap(arr, checkIdx, checkIdx+1)
					moved = true
				}
			}
			sortedNum ++
		} else {
			break
		}
	}
	return arr
}
