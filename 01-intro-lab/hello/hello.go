package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	s := "Hello from Golang!"
	fmt.Println(s)
	var goquote string = quote.Go()
	fmt.Println(goquote)
}
