package privatefunc

import "fmt"

func DoSomething(b bool, s string) {
	if b {
		return
	}

	if b == false {
		return
	}

	doPrivateThings(s)
}

func doPrivateThings(s string) {
	fmt.Println(s)
}