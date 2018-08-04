package main

import (
	"fmt"
	"math/rand"
)

// function description you can write here.
func bubble(sp *[]int) ([]int, error) {
	for i := 0; i < len(*sp); i += 1 {
		fmt.Println(*sp[i])
		*sp[i] = -*sp[i]
	}
	return *sp, nil
}

func main() {
	var Print = fmt.Println
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, rand.Intn(100))
	}
	Print(s)
	sortedS := bubble(&s)
	Print(sortedS)
}
