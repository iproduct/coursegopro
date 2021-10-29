package main

import (
	"fmt"
	"runtime"
	"time"
)
// Fake a long and difficult work.
func DoWork(n int, jobs <-chan struct{}) {
	fmt.Println("doing", n)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("finished", n)
	<-jobs // release the token
}
func main() {
	// concurrentJobs is a buffered channel implemting semaphore that blocks
	// if more than 20 goroutines are started at once
	var concurrentJobs = make(chan struct{}, 5)
	for i := 0; i < 30; i++ {
		concurrentJobs <- struct{}{} // acquire a  token
		go DoWork(i, concurrentJobs)
		fmt.Printf("Current number of goroutines: %d\n", runtime.NumGoroutine())
	}
}

