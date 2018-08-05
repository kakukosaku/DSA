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

func initSlice() {
	var s1 = []int{1, 2, 3}
	p(s1)
	fmt.Println(s1)
	s2 := make([]int, 3, 10)
	fmt.Println(s2)
	fmt.Println(s2[:5])
	var s3 []int
	p(s3)
	fmt.Println(s3)
	fmt.Println("s3 == nil", s3 == nil)
}

func sliceUsage(s []int) {
	for index, value := range s {
		s[index] = -value
	}
}

func testSliceUsage() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Just init a slice by explicity", s1)
	p(s1)
	s2 := append(s1, 6)
	p(s2)
	fmt.Println("append 6 to s1", s2)
	s2[0] = -s2[0]
	fmt.Println("after change s2 first ele, s1", s1)
	fmt.Println("after change s2 first ele, s2", s2)
	slice := make([]int, 5)
	fmt.Println("Just init a slice by make", slice)
	for i := 0; i < 5; i++ {
		slice[i] = i
	}
	fmt.Println("give value after for loop", slice)
	sliceUsage(slice)
	fmt.Println("after function invoke", slice)
}

func p(s []int) {
	fmt.Println("its len:", len(s))
	fmt.Println("its cap:", cap(s))
}

func main() {
	// fmt.Println("testSliceNegative")
	// testSliceNegative()
	// fmt.Println("testPointSliceNegative")
	// testPointSliceNegative()
	// fmt.Println("In vim-go ^_^")
	// fixArray()
	// initSlice()
	func() {
		fmt.Println("I am anonyous funcion")
	}()
	fmt.Println(10.0 / 3)
	if i, j := 0, 1; i < j {
		fmt.Println("valid")
	}
}
