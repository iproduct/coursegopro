package main

import(
	"fmt"
	"rsc.io/quote"
	"01-intro/stringutil"
)
var s = "Hello from Golang!"
func main() {
	fmt.Println(s)
	goquote := quote.Go()
	fmt.Println(goquote)
	fmt.Println(stringutil.Reverse(s))
}