package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	goDown := false
	restRows := make([]string, numRows)
	currRow := 0
	for _, v := range s {
		restRows[currRow] += string(v)
		if currRow == numRows-1 || currRow == 0 {
			goDown = !goDown
		}

		if goDown {
			currRow += 1
		} else {
			currRow -= 1
		}
	}

	return strings.Join(restRows, "")
}


func main() {
	fmt.Println(convert("ABCDE", 1))
	fmt.Println(convert("LEETCODEISHIRING", 3))
	// Output:
	// ABCDE
	// LCIRETOESIIGEDHN
}

