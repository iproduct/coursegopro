package main

import "fmt"

func main() {
	a := 1
	fmt.Printf("%p\n", &a)
	b, a := a, a
	//b++
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", &a)


	var a2 int
	var b2, c2 = &a2, &a2
	fmt.Println(b2, c2)   // 0x1040a124 0x1040a124
	fmt.Println(&b2, &c2) // 0x1040c108 0x1040c110
}
