// author: kaku
// date: 2019/月/日
//
// GitHub:
//   https://github.com/kakukosaku
//
// Description:
//
package main

func lengthOfLongestSubstring(s string) int {
	var recordMap map[string]int
	recordMap = make(map[string]int)
	maxNoRepeatSubStringLen := 0
	startIndex := 0
	for i, v := range s {
		valIdx, ok := recordMap[string(v)]
		if !ok || ok && (valIdx < startIndex) {
			recordMap[string(v)] = i
			if i-startIndex+1 > maxNoRepeatSubStringLen {
				maxNoRepeatSubStringLen = i - startIndex + 1
			}
		} else {
			startIndex = valIdx + 1
			recordMap[string(v)] = i
		}
	}

	return maxNoRepeatSubStringLen
}

func main() {
	rest := lengthOfLongestSubstring("abcabcbb")
	println(rest)
}
