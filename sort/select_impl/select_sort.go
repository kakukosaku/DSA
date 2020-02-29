// author: kaku
// date: 2019/02/27
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//
package selectimpl

import "github.com/kakukosaku/DSA/sort/utils"

func SelectSort(arr []int) []int {
	for sortedNum := 0; sortedNum < len(arr); sortedNum++ {
		currIdx := sortedNum
		smallIdx := sortedNum
		for checkIdx := sortedNum + 1; checkIdx < len(arr); checkIdx++ {
			// 从小到大
			if arr[checkIdx] < arr[smallIdx] {
				smallIdx = checkIdx
			}
		}

		utils.Swap(arr, currIdx, smallIdx)

	}
	return arr
}
