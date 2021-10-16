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
		fmt.Println("Zero value: <nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	//var i interface{}
	//describe(i) // (<nil>, <nil>)
	//
	//i = 42
	//describe(i) // (42, int)
	//
	//i = []interface{}{"hello"}
	//describe(i) // (hello, string)

	// (nil, T)
	var i2 I
	describe(i2)
	//i2.M()

	var t *T // zero value
	i2 = t
	describe(i2)
	i2.M()

	i2 = &T{"hello"}
	describe(i2)
	i2.M()
}
