package main

import "fmt"

type Result struct {
	value string
}

func deferDemo() (result1, result2 string) {
	defer func() {
		fmt.Printf("Inside defered func: %v, %v\n", result1, result2) // 2
		result1 = "NEW_VALUE"
	}()
	fmt.Printf("Inside deferDemo()\n") // 1
	result1 = "ABCD"
	// ... more code here ...
	result2 = "EFGH"
	return
}

func deferDemoPtr() (result1, result2 *string) {
	defer func() {
		fmt.Printf("Inside defered func: %v, %v\n", *result1, *result2) // 2
		*result1 = "NEW_VALUE"
	}()
	fmt.Printf("Inside deferDemoPtr()\n") // 1
	str1, str2 := "ABCD", "EFGH"
	return &str1, &str2
}

func main() {
	r1, r2 := deferDemo()
	fmt.Printf("Inside main(): %#v, %#v\n", r1, r2) // 3

	rp1, rp2 := deferDemoPtr()
	fmt.Printf("Inside main(): %#v, %#v\n", *rp1, *rp2) // 3

}
