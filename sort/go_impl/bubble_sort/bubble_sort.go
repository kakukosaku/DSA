package sort

import (
	"fmt"
	"math/rand"
)

// BubbleSort define: func bubble(sp []int) ([]int, error) {
func BubbleSort(s []int) []int {
	for i := 0; i < len(s); i++ {
		// 从 0 起, 从上往下冒泡...
		// 排好序的置于最后
		for j := 0; j < len(s)-1-i; j++ {
			// 从小到大
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			} else {
				continue
			}
		}
	}
	// don't need return slice, in this porpose.
	// we don't grow slice just change the corresponding array value.
	return s
}

func main() {
	// init a slice
	// var s []int
	// for i := 0; i < 10; i++ {
	// 	s = append(s, rand.Intn(100))
	// }

	// or you and do like this
	s := make([]int, 10, 10) // len 10, cap 10 later is optional.
	for index := range s {
		s[index] = rand.Intn(100)
	}
	fmt.Println("init slice:", s)
	sortedSlice := BubbleSort(s)
	fmt.Println("sort slice:", sortedSlice)
}
