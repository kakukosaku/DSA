package main

import "fmt"

func maxArea(arr []int) int {
	minIdx := 0
	maxIdx := len(arr) - 1
	var mostWater int

	for minIdx != maxIdx {
		_mostWater := min(arr[minIdx], arr[maxIdx]) * (maxIdx - minIdx)
		if _mostWater > mostWater {
			mostWater = _mostWater
		}

		if arr[minIdx] < arr[maxIdx] {
			minIdx++
			continue
		}

		maxIdx--
	}

	return mostWater
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(arr))
}
