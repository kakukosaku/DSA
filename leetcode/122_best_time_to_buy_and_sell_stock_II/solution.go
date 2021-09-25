package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxGain := 0
	for i, price := range prices {
		if i == 0 {
			continue
		}
		maxGain += max(0, price - prices[i-1])
	}

	return maxGain
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
