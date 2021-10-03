package main

import (
	"fmt"
	"rsc.io/quote"
	"github.com/iproduct/coursegopro/01-intro-lab/stringutil"
)

func main() {
	s := "Hello from Golang!"
	fmt.Println(s)
	var goquote string = quote.Go()
	fmt.Println(goquote)
	fmt.Println(stringutil.Reverse(goquote))	
}
