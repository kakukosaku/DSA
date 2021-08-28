package main

import "fmt"

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) == 0 {
		return result
	}

	var (
		candidate = make([]int, len(nums))
		currIdx   int
	)
	copy(candidate, nums)
	result = backtrack(candidate, result, currIdx)
	return result
}

func backtrack(candidate []int, result [][]int, currIdx int) [][]int {
	if currIdx == len(candidate)-1 {
		tmp := make([]int, len(candidate))
		copy(tmp, candidate)
		result = append(result, tmp)
		return result
	}

	for i := currIdx; i < len(candidate); i++ {
		swap(candidate, currIdx, i)
		result = backtrack(candidate, result, currIdx+1)
		swap(candidate, currIdx, i)
	}

	return result
}

func swap(s []int, a, b int) {
	s[a], s[b] = s[b], s[a]
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
