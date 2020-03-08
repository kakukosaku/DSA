// author: kaku
// date: 2020/03/09
//
// GitHub:
//	https://github.com/kakukosaku
//
// Description:
//
package main

import "fmt"

func solution(n int) int {
	totalPoint := 10
	if n == 1 {
		return 0
	}

	// dp[i][j] 表示, 从i点, 走j步, 回到原点的方案
	dp := make([][]int, totalPoint)
	for i := 0; i < totalPoint; i++ {
		dp[i] = make([]int, totalPoint)
	}
	dp[0][0] = 1

	for j := 1; j <= n; j++ {
		for i := 0; i < totalPoint; i++ {
			leftPoint := i - 1
			rightPoint := i + 1
			if leftPoint < 0 {
				leftPoint += totalPoint
			}
			if rightPoint >= totalPoint {
				rightPoint -= totalPoint
			}
			dp[i][j] = dp[leftPoint][j-1] + dp[rightPoint][j-1]
		}
	}
	for _, v := range dp {
		fmt.Printf("%v\n", v)
	}
	return dp[0][n]
}

func main() {
	fmt.Println(solution(2))
	fmt.Println(solution(3))
	fmt.Println(solution(4))
}
