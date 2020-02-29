// author: kaku
// date: 2019/02/27
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//
package insert

func InsertSort(arr []int) []int {
	for sortedNum := 0; sortedNum < len(arr); sortedNum ++ {
		checkIdx := sortedNum + 1
		// From small to large
		if checkIdx < len(arr) && arr[checkIdx] < arr[checkIdx - 1] {
			temp := arr[checkIdx]

			for ; checkIdx > 0 && arr[checkIdx - 1] > temp; checkIdx -- {
				// move check_elem to next position
				arr[checkIdx] = arr[checkIdx - 1]
			}
			arr[checkIdx] = temp
		}
	}

	return arr
}
