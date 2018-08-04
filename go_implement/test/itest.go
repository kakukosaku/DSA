package main

import "fmt"

func testSliceNegative() {
	s := []int{1, 2, 3}
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		s[i] = -s[i]
	}
	fmt.Println(s)
}

func testPointSliceNegative() {
	var sp *[]int
	s := []int{1, 2, 3}
	sp = &s
	*sp = append(*sp, 4)
	fmt.Println("s", s)
	fmt.Println("sp", sp)
	for i := 0; i < len(*sp); i++ {
		(*sp)[i] = -(*sp)[i]
	}
	fmt.Println("slice", s)
	fmt.Println("slice point", sp)
}

func main() {
	fmt.Println("testSliceNegative")
	testSliceNegative()
	fmt.Println("testPointSliceNegative")
	fmt.Println("In vim-go ^_^")
}
