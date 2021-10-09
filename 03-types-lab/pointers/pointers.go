package main

import "fmt"

func main() {
	a := []int{1}
	a2 := a
	fmt.Printf("%p, %p", &a2[0], &a[0])
}
