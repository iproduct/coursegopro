package main

import "fmt"

func main() {
	ch := make(chan string)
	go func(output chan string) {
		for i := 0; i < 10; i++ {
			output <- fmt.Sprintf("sending N=%d", i)
		}
		close(output)
	}(ch)
	var val string
	for ok := true; ok; {
		val, ok = <-ch
		fmt.Printf("Received: %#v, %#v\n", val, ok)
	}
	ch <- fmt.Sprintf("where is your towel?")
}
