package main

import(
	"fmt"
	"rsc.io/quote"
	"01-intro/stringutil"
)

func main() {
	s := "Hello from Golang!"
	fmt.Println(s)
	goquote := quote.Go()
	fmt.Println(goquote)
	fmt.Println(stringutil.Reverse(s))
}