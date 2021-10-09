package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "Golang"
	for _, s := range a {
		fmt.Printf("%#v \n", s)
	}
	fmt.Printf("%#v, len = %d, cap = %d\n", a, len(a), cap(a))

}
