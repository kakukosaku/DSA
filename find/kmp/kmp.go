// author: kaku
// date: 2020/03/09
//
// GitHub:
//	https://github.com/kakukosaku
//
// Description:
//
package main

import (
	"fmt"
)

func GetNext(t string) []int {
	fmt.Printf("Target string:\n\t%v\n", t)
	next := make([]int, len(t))
	i := 0
	k := -1
	next[0] = -1
	for i < len(t)-1 {
		if k == -1 || t[i:i+1] == t[k:k+1] {
			k++
			i++
			next[i] = k
			fmt.Printf("第 %d 个元素的 *next* 数组:\t%v\n", i+1, next[:i+1])
		} else {
			k = next[k]
		}
	}

	return next
}

func KMPMatch(s, t string) bool {
	next := GetNext(t)
	startIdx := 0
	patternStartIdx := 0

	for startIdx < len(s) && patternStartIdx < len(t) {
		if patternStartIdx == -1 || s[startIdx] == t[patternStartIdx] {
			startIdx++
			patternStartIdx++
		} else {
			patternStartIdx = next[patternStartIdx]
		}
	}

	if patternStartIdx == len(t) {
		return true
	}

	return false
}

func main() {
	println(KMPMatch("bbc abcdab abcdabcdabde", "abcdabd"))
}
