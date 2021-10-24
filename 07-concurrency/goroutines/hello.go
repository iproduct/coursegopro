package main

import (
	"fmt"
	"time"
)

func say(msg string) {
	fmt.Println(msg)
}

func hello() {
	go say("hello")
	say("world")

	time.Sleep(1 * time.Second)
}
