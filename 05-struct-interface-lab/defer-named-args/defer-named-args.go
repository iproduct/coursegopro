package main

import "fmt"

type Result struct {
	value string
}

func deferDemo() (result1, result2 int) {
	defer func() {
		fmt.Printf("Inside defered func: %v, %v\n", result1, result2) // 2
		result1 = 108
	}()
	fmt.Printf("Inside deferDemo()\n") // 1
	result1 = 2
	// ... more code here ...
	result2 = 3
	return
}

func main() {
	r1, r2 := deferDemo()
	fmt.Printf("Inside main(): %#v, %#v\n", r1, r2); // 3
}
