package main

import "fmt"
type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i interface{}
	describe(i) // (<nil>, <nil>)

	i = 42
	describe(i) // (42, int)

	i = "hello"
	describe(i) // (hello, string)

	// (nil, T)
	var i2 I

	var t *T
	i2 = t
	describe(i2)
	i2.M()

	i2 = &T{"hello"}
	describe(i2)
	i2.M()
}


