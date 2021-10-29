package main

import "fmt"

func main3() {
	var c chan string
	c <- "message" // Deadlock
}
func main2() {
	var c chan string
	fmt.Println(<-c) // Deadlock
}

func main() {
	select {}
	fmt.Println("Demo finished")
}