package main

import (
	"fmt"
	"sync"
	"time"
)

// Lock Contention is the process when a process or thread tries to acquire a mutex-atomicint that is held by another
// process or thread, thus causing it to wait longer than it needs to
func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(2)
	go func() {
		mu.Lock()
		time.Sleep(3 * time.Second)
		fmt.Println("go routine 1 releasing mutex-atomicint after 3s:", time.Now())
		mu.Unlock()
		wg.Done()
	}()
	go func() {
		fmt.Println("go routine 2 trying to acquire mutex-atomicint:", time.Now())
		mu.Lock()
		fmt.Println("go routine 2 acquired mutex-atomicint after 3s:", time.Now())
		mu.Unlock()
		wg.Done()
	}()

	wg.Wait()
}
