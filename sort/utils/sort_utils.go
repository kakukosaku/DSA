// author: kaku
// date: 2019/02/27
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//

package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomArray get random array.
func RandomArray(size int) []int {
	rand.Seed(time.Now().UTC().UnixNano())

	arr := make([]int, 0)
	for i := 0; i < size; i++ {
		arr = append(arr, rand.Intn(100))
	}

	return arr
}


// Swap swap i and j
func Swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}


// Show show array
func Show(arr []int, description string) {
	fmt.Printf("\nDescription:\t%s\n", description)
	fmt.Printf("Show array:\t%v\n", arr)
}


// Check array is sorted or not
func IsSorted(arr []int, desc bool) bool {
	var sort bool
	for i := 0; i < len(arr) - 1; i++ {
		if desc {
			sort = arr[i] >= arr[i+1]
		} else {
			sort = arr[i] <= arr[i+1]
		}

		if !sort {
			fmt.Println("* Result:\tNot sorted array!")
			return sort
		}
	}

	fmt.Println("* Result:\tSorted array!")
	return sort
}
