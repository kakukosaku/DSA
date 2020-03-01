// author: kaku
// date: 2020/02/29
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//		leetcode #1143 https://leetcode-cn.com/problems/longest-common-subsequence/
package main

import (
	"fmt"
)

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}

	t1s := []byte(text1)
	t2s := []byte(text2)
	dp := make([][]int, len(t1s)+1)
	// 初始化dp数组
	for i := range dp {
		dp[i] = make([]int, len(t2s)+1)
		println(i)
	}

	f := func(i, j int) int {
		var m int
		if j > i {
			m = j
		} else {
			m = i
		}
		return m
	}

	for i := 1; i < len(t1s)+1; i++ {
		for j := 1; j < len(t2s)+1; j++ {
			if t1s[i-1] == t2s[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = f(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	for _, v := range dp {
		fmt.Printf("%v\n", v)
	}

	return dp[len(t1s)][len(t2s)]
}

func main() {
	text1 := "abcde"
	text2 := "ace"
	fmt.Println(longestCommonSubsequence(text1, text2))
}
