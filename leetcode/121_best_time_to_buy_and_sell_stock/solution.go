package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	max := 0
	minPrice := prices[0]
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}

		if price-minPrice > max {
			max = price - minPrice
		}
	}

	return max
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
