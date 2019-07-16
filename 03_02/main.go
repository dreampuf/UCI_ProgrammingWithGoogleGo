package main

import (
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	/*
	   Following instructions introducing a race condition case
	   Both goroutines obtain the variable `x` and operating its relative synchronous.
	   Because the interleavings, goroutine 1 may be executed after goroutine 2, or before goroutine 2, the `x` in goroutine 1  may 1 or 2, either same situation for goroutine 2.
	 */

	x := 0
	wg.Add(1)
	go func() {
		x += 1
		println(x)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		x += 1
		println(x)
		wg.Done()
	}()

	wg.Wait()
}
