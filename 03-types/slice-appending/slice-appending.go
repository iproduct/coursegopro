package main

import "fmt"

func main() {
	//var s []int
	//printSlice(s) // len=0 cap=0 []
	//var s2 []int = make([]int, 0)
	//printSlice(s2) // len=0 cap=0 []

	// append works on nil slices.
	//s2 := append(s, 42)
	//printSlice(s2) // len=1 cap=1 [0]
	//fmt.Printf("%p, %p\n", s, s2)
	//
	//s3 := append(s2, 43)
	//printSlice(s3) // len=1 cap=1 [0]
	//fmt.Printf("%p, %p\n", s3, s2)
	//
	//s4 := append(s3, 44)
	//printSlice(s4) // len=1 cap=1 [0]
	//fmt.Printf("%p, %p\n", s4, s3)
	//
	//s5 := append(s4, 45)
	//printSlice(s5) // len=1 cap=1 [0]
	//fmt.Printf("%p, %p\n", s5, s4)
	//
	//// The slice grows as needed.
	//s6 := append(s2, 1)
	//printSlice(s6)                                   // len=2 cap=2 [0 1]
	//fmt.Printf("Same array: %t\n", &s6[0] == &s2[0]) // Same array: false

	//// The array is copied only if needed
	a := [...]int{2, 3, 5, 7, 9}
	s4 := a[1:3]
	printSlice(s4) // len=2 cap=4 [3 5]
	s5 := append(s4, 11, 13)
	printSlice(s5)                                   // len=4 cap=4 [3 5 11 13]
	fmt.Printf("Same array: %t\n", &s5[0] == &s4[0]) // Same array: true
	s6 := append(s5, 17)
	printSlice(s6)                                   // len=5 cap=8 [3 5 11 13 17]
	fmt.Printf("Same array: %t\n", &s6[0] == &s5[0]) // // Same array: false

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %#v\n", len(s), cap(s), s)
}
