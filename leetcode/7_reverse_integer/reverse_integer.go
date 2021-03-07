package main

import "math"

func reverse(x int) int {
	var remind, result int
	for x != 0 {
		remind = x % 10
		x = x / 10
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && remind > 7) {
			return 0
		}
		if result < math.MinInt32/10 || (result == math.MinInt32/10 && remind < -8) {
			return 0
		}
		result = result*10 + remind
	}
	return result
}

func main() {
	println(reverse(123))
	println(reverse(-123))
}
